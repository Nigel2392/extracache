package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Message struct {
	Channel_ID int         `json:"channel_id"`
	Type       string      `json:"type"`
	Config     *Config     `json:"config"`
	Key        string      `json:"key"`
	Val        interface{} `json:"val"`
	TTL        int         `json:"ttl"`
}

type Server struct {
	Addr string `json:"addr"`
	Port int    `json:"port"`
}

func (s *Server) Str_port() string {
	return strconv.Itoa(s.Port)
}

func (s *Server) Listen() {
	LOGGER.Info("Listening on " + s.Addr + ":" + s.Str_port())
	ln, err := net.Listen("tcp", s.Addr+":"+s.Str_port())
	if err != nil {
		err = errors.New("error creating server: " + err.Error())
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			err = errors.New("error accepting connection: " + err.Error())
			LOGGER.Error(err.Error())
			break
		}
		LOGGER.Info("Accepted connection from " + conn.RemoteAddr().String())
		go HandleConnection(conn)
	}
}

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 8192)
		data, err := conn.Read(buf)
		if err != nil {
			LOGGER.Error("Error reading from connection: " + err.Error())
			break
		}
		bytes := buf[:data]
		LOGGER.Debug("Received data: " + string(bytes))
		var msg Message
		json.Unmarshal(bytes, &msg)
		fmt.Println(msg)
		data_to_json := ParseRequest(msg)
		if json_data, err := json.Marshal(data_to_json); err == nil {
			conn.Write(json_data)
			LOGGER.Info("Sent data to: " + conn.RemoteAddr().String())
		} else {
			break
		}
	}
}

func ParseRequest(msg Message) map[string]interface{} {
	data_to_json := make(map[string]any)
	switch strings.ToUpper(msg.Type) {
	case "SET":
		LOGGER.Info("Received message from socket: " + msg.Type)
		// fmt.Println(msg.Type)
		CACHE_LIST.SetItemInCache(msg.Channel_ID, msg.Key, msg.Val, msg.TTL)
		data_to_json["DATA"] = map[string]any{"KEY": msg.Key}
		data_to_json["CHANNEL"] = msg.Channel_ID
		data_to_json["STATUS"] = "OK"

	case "GET":
		LOGGER.Info("Received message from socket: " + msg.Type)
		// fmt.Println(msg.Type)
		data := CACHE_LIST.GetItemFromCache(msg.Channel_ID, msg.Key)
		data_to_json["DATA"] = map[string]any{"KEY": msg.Key, "VALUE": data}
		data_to_json["CHANNEL"] = msg.Channel_ID
		data_to_json["STATUS"] = "OK"
		if data == nil {
			data_to_json["STATUS"] = "NOT_FOUND"
		}

	case "HASKEY":
		LOGGER.Info("Received message from socket: " + msg.Type)
		// fmt.Println(msg.Type)
		data := CACHE_LIST.CacheHasKey(msg.Channel_ID, msg.Key)
		data_to_json["DATA"] = map[string]any{"KEY": msg.Key, "VALUE": data}
		data_to_json["CHANNEL"] = msg.Channel_ID
		data_to_json["STATUS"] = "OK"
		if !data {
			data_to_json["STATUS"] = "NOT_FOUND"
		}

	case "DELETE":
		LOGGER.Info("Received message from socket: " + msg.Type)
		// fmt.Println(msg.Type)
		success := CACHE_LIST.DeleteItemFromCache(msg.Channel_ID, msg.Key)
		data_to_json["DATA"] = map[string]any{"KEY": msg.Key, "SUCCESS": success}
		data_to_json["CHANNEL"] = msg.Channel_ID
		data_to_json["STATUS"] = "OK"
		if !success {
			data_to_json["STATUS"] = "NOT_FOUND"
		}

	case "SIZE":
		LOGGER.Info("Received message from socket: " + msg.Type)
		// fmt.Println(msg.Type)
		size := CACHE_LIST.GetCacheSize(msg.Channel_ID)
		data_to_json["DATA"] = map[string]any{"SIZE": size}
		data_to_json["CHANNEL"] = msg.Channel_ID
		data_to_json["STATUS"] = "OK"

	case "SIZEALL":
		LOGGER.Info("Received message from socket: " + msg.Type)
		// fmt.Println(msg.Type)
		data_to_json["DATA"] = map[string]any{"SIZE_ALL": CACHE_LIST.GetCacheSizeAll()}
		data_to_json["STATUS"] = "OK"

	case "KEYS":
		LOGGER.Info("Received message from socket: " + msg.Type)
		// fmt.Println(msg.Type)
		keys := CACHE_LIST.GetCacheKeys(msg.Channel_ID)
		data_to_json["DATA"] = map[string]any{"KEYS": keys}
		data_to_json["CHANNEL"] = msg.Channel_ID
		data_to_json["STATUS"] = "OK"

	case "CONFIGURE":
		LOGGER.Info("Received message from socket: " + msg.Type)
		// fmt.Println(msg.Type)
		config := msg.Config
		if config != nil {
			config_json, err := json.MarshalIndent(config, "", " ")
			if err != nil {
				data_to_json["DATA"] = map[string]any{"ERROR": err.Error()}
				data_to_json["STATUS"] = "ERROR"
				data_to_json["TYPE"] = "configure"
			}
			err = CONF.WriteConfig(config_json)
			if err != nil {
				data_to_json["DATA"] = map[string]any{"ERROR": err.Error()}
				data_to_json["STATUS"] = "ERROR"
			} else {
				data_to_json["DATA"] = map[string]any{"CONFIG": config}
				data_to_json["STATUS"] = "OK"
			}
			data_to_json["TYPE"] = "configure"
		}
	// No cases were matched
	default:
		LOGGER.Warning("Received message from socket: " + msg.Type)
		data_to_json["DATA"] = map[string]any{"ERROR": "Invalid request"}
		data_to_json["STATUS"] = "ERROR"
	}
	data_to_json["TYPE"] = msg.Type
	return data_to_json
}
