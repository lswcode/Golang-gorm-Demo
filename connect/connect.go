package connect

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	DataBase string `json:"database"`
	LogoMode bool   `json:"logo_mode"`
}

func Connect() *MysqlConfig {
	mysqlConfig := MysqlConfig{}
	file, err := os.Open("./mysql.json") // 打开json文件
	if err != nil {
		panic(err)
	}

	defer file.Close()

	byteData, err2 := ioutil.ReadAll(file)

	if err2 != nil {
		panic(err2)
	}

	err3 := json.Unmarshal(byteData, &mysqlConfig)
	if err3 != nil {
		panic(err3)
	}

	return &mysqlConfig
}
