package configuration

//MischiefConfig is for configuring a chaos command
type MischiefConfig struct {
	TargetNamespace string
}

//NewDefaultConfiguration provides a default configuration that will delete pods
func NewDefaultConfiguration() *MischiefConfig {
	return &MischiefConfig{}
}
