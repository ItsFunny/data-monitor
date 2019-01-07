package capture

import (
	"github.com/Sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var log = logrus.New()

type CaptureConfig struct {
	ifaces []string // 绑定的设备名称

}

func (c CaptureConfig) validate() error {
	return nil
}

func LoadConfig(configPath string) *CaptureConfig {
	config := &CaptureConfig{}
	bytes, e := ioutil.ReadFile(configPath)
	if nil != e {
		log.Error("[LoadConfig]occur error:%v", e)
		os.Exit(-1)
	}
	e = yaml.Unmarshal(bytes, config)
	if nil != e {
		log.Error("[yaml#unmarshal]occur error:%v", e)
		os.Exit(-1)
	}
	e = config.validate()
	if nil != e {
		log.Error("[CatureConfig#validate]occur error:%v", e)
		os.Exit(-1)
	}
	return config
}
