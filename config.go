package wxcoresdk

type ConfigStruct struct {
	GoCore goCoreConfig `yaml:"goCore"`
}

type goCoreConfig struct {
	Url string `yaml:"url"`
	Api struct {
		SendMsg string `yaml:"sendMsg"`
	} `yaml:"api"`
}

var Config *ConfigStruct

const (
	SendTypeText = 1
	SendTypeLink = 2
	SendTypePic  = 4
)
