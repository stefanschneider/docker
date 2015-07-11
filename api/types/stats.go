// This package is used for API stability in the types and response to the
// consumers of the API stats endpoint.
package types

type CommonCpuStats struct {
	CpuUsage CpuUsage `json:"cpu_usage"`
}

// All CPU stats are aggregated since container inception.
type CommonCpuUsage struct {
	// Total CPU time consumed.
	// Units: nanoseconds.
	TotalUsage uint64 `json:"total_usage"`
}

type CommonMemoryStats struct {
	// current res_counter usage for memory
	Usage uint64 `json:"usage"`
}

type CommonStats struct {
	CpuStats    CpuStats    `json:"cpu_stats,omitempty"`
	MemoryStats MemoryStats `json:"memory_stats,omitempty"`
}
