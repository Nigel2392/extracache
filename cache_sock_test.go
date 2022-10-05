package main

import (
	"errors"
	"net"
	"testing"
)

func Test_Listen(t *testing.T) {
	var srv Server = Server{Addr: "127.0.0.1", Port: 3882}
	go srv.Listen()
	conn, err := net.Dial("tcp", "127.0.0.1:3882")
	if err != nil {
		t.Errorf("Cannot connect to server")
	}
	conn.Close()
	LOGGER.Test("SOCKET Test_Listen Finished")
}
func Test_HandleConnection(t *testing.T) {
	var ln net.Listener
	var err error
	go func() {
		ln, err = net.Listen("tcp", "127.0.0.1:3883")
		if err != nil {
			t.Errorf("Cannot listen on port")
		}
		conn, err := ln.Accept()
		if err != nil {
			err = errors.New("error accepting connection: " + err.Error())
			t.Errorf(err.Error())
		}
		LOGGER.Info("Accepted connection from " + conn.RemoteAddr().String())
		HandleConnection(conn)
	}()
	conn, err := net.Dial("tcp", "127.0.0.1:3883")
	if err != nil {
		t.Errorf("Cannot connect to server")
	}
	conn.Close()
	LOGGER.Test("SOCKET Test_HandleConnection Finished")
}
func Test_ParseRequest(t *testing.T) {
	cache := Cache{Channel_ID: 1, Data: make(map[string]*CachedItem)}
	CACHE_LIST = &CacheList{Caches: make(map[int]*Cache)}
	CACHE_LIST.Caches[1] = &cache
	var msg Message = Message{Type: "SET", Key: "key", Val: "value", TTL: 10, Channel_ID: 1}
	data_to_json := ParseRequest(msg)
	if data_to_json["STATUS"] != "OK" {
		t.Errorf("Wrong status")
	}
	LOGGER.Test("SOCKET Test_ParseRequest Finished")
}
