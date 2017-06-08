package configuration

//MischiefConfig is for configuring a chaos command
type MischiefConfig struct {
	TargetNamespace string
	AttackCount     int
}

//NewDefaultConfiguration provides a default configuration that will delete pods
func NewDefaultConfiguration() *MischiefConfig {
	return &MischiefConfig{}
}
