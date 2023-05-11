package core

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"web/config"
	"web/global"

	"gopkg.in/yaml.v3"
)
const ConfigFile = "settings.yaml"

func InitConf(){
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success.")
	global.Config = c
}

func SetYaml() error{
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		// global.Log.Error(err)
		return err
	}
	err = ioutil.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		// global.Log.Error(err)
		return err
	}
	global.Log.Info("配置文件修改成功")
	return nil
}
