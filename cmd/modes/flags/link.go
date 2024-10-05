package flags

type LinkFlags struct {
	LinkLen int    `mapstructure:"tail_length"`
	Domain  string `mapstructure:"domain"`
}
