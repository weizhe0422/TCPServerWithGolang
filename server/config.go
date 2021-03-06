package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	ConnectMethod      string `json:"connectMethod"`
	ServerAddress      string `json:"serverAddress"`
	SocketPort         int    `json:"socketPort"`
	HttpPort           int    `json:"httpPort"`
	ServerStatusPath   string `json:"serverStatusPath"`
	ReceiveBuffer      int    `json:"receiveBuffer"`
	ApiSvrReadTimeOut  int    `json:"apiSvrReadTimeOut"`
	ApiSvrWriteTimeOut int    `json:"apiSvrWriteTimeOut"`
	RateLimitPerSecond int    `json:"rateLimitPerSecond"`
	RateLimitBuffer    int    `json:"rateLimitBuffer"`
	WebRoot            string `json:"webRoot"`
}

var (
	G_Config *Config
)

func InitConfig(fileName string) (err error) {
	var (
		content []byte
		conf    Config
	)

	if content, err = ioutil.ReadFile(fileName); err != nil {
		log.Println("failed to read configuration file: ", fileName, ", ", err.Error())
		return
	}

	if err = json.Unmarshal(content, &conf); err != nil {
		log.Println("failed to  Unmarshal configuration file:", err.Error())
		return
	}

	G_Config = &conf

	log.Println(G_Config)
	return
}
