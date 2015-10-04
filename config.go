package main

import (
    "encoding/json"
    "log"
    "os"
)

type Config struct {
    ListenAddr string
    Storage string
}

var conf Config

func LoadConf(){
	r, err := os.Open("config.json")
    if err != nil {
        log.Fatalln(err)
    }
    decoder := json.NewDecoder(r)
    err = decoder.Decode(&conf)
    if err != nil {
        log.Fatalln(err)
    }
}