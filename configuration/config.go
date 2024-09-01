package configuration

type GlobalConfiguration struct {
	Host   string
	APIKey string
}

func GetConfiguration() GlobalConfiguration {
	gc := &GlobalConfiguration{}

	gc.Host = ""   // TODO poke api host here
	gc.APIKey = "" // TODO do we need an API?

	return *gc
}
