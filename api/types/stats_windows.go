// This package is used for API stability in the types and response to the
// consumers of the API stats endpoint.
package types

type CpuStats struct {
	CommonCpuStats

	// Below here are platform-specific CPU stats
}

// All CPU stats are aggregated since container inception.
type CpuUsage struct {
	CommonCpuUsage

	// Below here are platform specific CPU usage stats
}

type MemoryStats struct {
	CommonMemoryStats

	// Below here are platform-specfic memory stats
}

type Stats struct {
	CommonStats

	// Below here are platform-specific stats
}
