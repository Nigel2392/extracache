package main

import (
	"ExtraCache/extralogger"
	"encoding/json"
	"errors"
	"os"
)

type Config struct {
	CACHE_CHANNELS int                 `json:"cache_channels"`
	CACHE_MAX_TTL  int                 `json:"cache_max_ttl"`
	CACHE_SAVE     bool                `json:"cache_save"`
	SAVE_FILE_NAME string              `json:"save_file_name"`
	SERVER         *Server             `json:"server"`
	LOGGER         *extralogger.Logger `json:"logger"`
}

func ParseConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create("config.json")
			if err != nil {
				err = errors.New("could not create config file")
				panic(err)
			}
			_, err = file.WriteString(`{
 "cache_channels":3,
 "cache_max_ttl":84600,
 "cache_save":true,
 "save_file_name":"LAST_CACHE",
 "server":{
  "addr":"127.0.0.1",
  "port":3239
 },
 "logger":{
  "level":"debug",
  "use_file":true,
  "path":".\\logs",
  "filename":"extracache.log",
  "max_lines":1000
 }
}`)
			if err != nil {
				err = errors.New("could not write to config file")
				panic(err)
			}
			file.Close()
			file, err = os.Open("config.json")
			if err != nil {
				err = errors.New("could not open config file")
				panic(err)
			}
		} else {
			err = errors.New("could not open config file")
			panic(err)
		}
	}
	defer file.Close()
	// Set up config
	CONF.SERVER = SERVER
	CONF.LOGGER = LOGGER
	// Decode json data into struct
	json.NewDecoder(file).Decode(&CONF)
	LOGGER.SetupFile()
}

func (c *Config) WriteConfig(data []byte) error {
	file, err := os.OpenFile("config.json", os.O_WRONLY, 0644)
	if err != nil {
		err = errors.New("could not open config file")
		return err
	}
	defer file.Close()
	file.Write(data)
	return nil
}
