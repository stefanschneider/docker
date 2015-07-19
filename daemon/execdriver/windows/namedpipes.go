// +build windows

package windows

import (
	"io"

	"github.com/Sirupsen/logrus"
	"github.com/docker/docker/daemon/execdriver"
)

// General comment. Handling I/O for a container is very different to Linux.
// We use a named pipe to HCS to copy I/O both in and out of the container,
// very similar to how docker daemon communicates with a CLI.

// startStdinCopy asynchronously copies an io.Reader to the container's
// process's stdin pipe and closes the pipe when there is no more data to copy.
func startStdinCopy(pipe io.WriteCloser, copyfrom io.Reader) {

	// Anything that comes from the client stdin should be copied
	// across to the stdin named pipe of the container.
	if copyfrom != nil {
		go func() {
			defer pipe.Close()
			logrus.Debugln("Calling io.Copy on stdin")
			bytes, err := io.Copy(pipe, copyfrom)
			logrus.Debugf("Finished io.Copy on stdin bytes=%d err=%s", bytes, err)
		}()
	} else {
		pipe.Close()
	}
}

// startStdouterrCopy asynchronously copies data from the container's process's
// stdout or stderr pipe to an io.Writer and closes the pipe when there is no
// more data to copy.
func startStdouterrCopy(pipe io.ReadCloser, pipeName string, copyto io.Writer) {
	// Anything that comes from the container named pipe stdout/err should be copied
	// across to the stdout/err of the client
	if copyto != nil {
		go func() {
			defer pipe.Close()
			logrus.Debugln("Calling io.Copy on", pipeName)
			bytes, err := io.Copy(copyto, pipe)
			logrus.Debugf("Copied %d bytes from %s", bytes, pipeName)
			if err != nil {
				// Not fatal, just debug log it
				logrus.Debugf("Error hit during copy %s", err)
			}
		}()
	} else {
		pipe.Close()
	}
}

// setupPipes starts the asynchronous copying of data to and from the named
// pipes used byt he HCS for the std handles.
func setupPipes(stdin io.WriteCloser, stdout, stderr io.ReadCloser, pipes *execdriver.Pipes) {
	if stdin != nil {
		startStdinCopy(stdin, pipes.Stdin)
	}
	if stdout != nil {
		startStdouterrCopy(stdout, "stdout", pipes.Stdout)
	}
	if stderr != nil {
		startStdouterrCopy(stderr, "stderr", pipes.Stderr)
	}
}
