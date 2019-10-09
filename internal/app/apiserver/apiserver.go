package apiserver

type apiserver struct {
	config *Config
}

func New(config *Config) *apiserver {
	return &apiserver{
		config: config,
	}
}

// start
func (s *apiserver) Start() error {
	return nil
}
