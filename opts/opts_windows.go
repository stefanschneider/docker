// +build windows

package opts

// On Windows, docker by default listens on a named pipe
var DefaultLocalAddr = `npipe:\\.\pipe\docker-daemon`
