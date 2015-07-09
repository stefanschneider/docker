package runconfig

// ContainerConfigWrapper is a Config wrapper that hold the container Config (portable)
// and the corresponding HostConfig (non-portable).
type ContainerConfigWrapper struct {
	*Config
	HostConfig *HostConfig `json:"HostConfig,omitempty"`
}

// GetHostConfig gets the HostConfig of the Config.
func (w *ContainerConfigWrapper) GetHostConfig() *HostConfig {
	return w.HostConfig
}
