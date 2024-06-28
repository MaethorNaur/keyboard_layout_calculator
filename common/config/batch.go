package config

type editAction func(c *Config) bool

type batchEdit struct {
	actions []editAction
	config  *Config
}

func (be *batchEdit) Add(action editAction) *batchEdit {
	be.actions = append(be.actions, action)
	return be
}

func (be *batchEdit) Execute() (edited bool, err error) {
	edited = false
	for _, cb := range be.actions {
		if cb(be.config) {
			edited = true
		}
	}
	if edited {
		err = be.config.Save()
	}

	return
}
