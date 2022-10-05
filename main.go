package main

import (
	"ExtraCache/extralogger"
	"ExtraCache/typeutils"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/TwiN/go-color"
)

var MUTEX *sync.Mutex
var MUTEX_WATCHER *sync.Mutex
var CONF Config
var CACHE_LIST *CacheList
var SERVER *Server
var LOGGER *extralogger.Logger

func CacheWatcher() {
	for {
		time.Sleep(1 * time.Second)
		for _, cache := range CACHE_LIST.Caches {
			go func(cache *Cache) {
				MUTEX_WATCHER.Lock()
				for i, cached_item := range cache.Data {
					if cached_item.IsExpired() {
						cache.Delete(cached_item.Key)
					} else {
						cached_item.TTL -= 1
						cache.Data[i] = cached_item
					}
				}
				MUTEX_WATCHER.Unlock()
			}(cache)
		}
	}
}

func init() {
	// Set up cache mutex
	MUTEX = &sync.Mutex{}
	// Set up watcher mutex
	MUTEX_WATCHER = &sync.Mutex{}
	// Set up server
	SERVER = &Server{}
	// Set up logger
	LOGGER = &extralogger.Logger{}

	// Parse config
	ParseConfig()
	// Print logo
	PrintLogo()
	// Print config
	PrintOptions()
	// Check if cache items should expire.
	go CacheWatcher()
}

func main() {
	if len(os.Args) > 1 {
		LOGGER.Error("This program doesn't take any arguments.")
	}
	CACHE_LIST = &CacheList{make(map[int]*Cache)}
	for i := 0; i < CONF.CACHE_CHANNELS; i++ {
		Cache := Cache{Channel_ID: i}
		Cache.Data = make(map[string]*CachedItem, 512)
		CACHE_LIST.Caches[i] = &Cache
	}

	SERVER.Listen()
}

func PrintLogo() {
	// fmt.Println(color.Colorize(color.Purple, "#"+typeutils.Repeat("#", 70)))
	fmt.Println(color.Colorize(color.Purple, `
	
███████╗██╗  ██╗████████╗██████╗  █████╗    ██████╗ █████╗  ██████╗██╗  ██╗███████╗
██╔════╝╚██╗██╔╝╚══██╔══╝██╔══██╗██╔══██╗  ██╔════╝██╔══██╗██╔════╝██║  ██║██╔════╝
█████╗   ╚███╔╝    ██║   ██████╔╝███████║  ██║     ███████║██║     ███████║█████╗  
██╔══╝   ██╔██╗    ██║   ██╔══██╗██╔══██║  ██║     ██╔══██║██║     ██╔══██║██╔══╝  
███████╗██╔╝ ██╗   ██║   ██║  ██║██║  ██║  ╚██████╗██║  ██║╚██████╗██║  ██║███████╗
╚══════╝╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝   ╚═════╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝╚══════╝

`+color.Colorize(color.Red, `Server: V 1.0.0`)+`
`+color.Colorize(color.Blue, `© Nigel van Keulen - ITExtra - 2022`)+`

	`))
	fmt.Println(color.Colorize(color.Purple, "#"+typeutils.Repeat("#", 70)))
}

func PrintOptions() {
	LOGGER.Info("Creating " + strconv.Itoa(CONF.CACHE_CHANNELS) + " channels")
	LOGGER.Info("Max TTL: " + strconv.Itoa(CONF.CACHE_MAX_TTL))
	LOGGER.Info("Save: " + strconv.FormatBool(CONF.CACHE_SAVE))
	fmt.Println(typeutils.Repeat("-", 70))
	LOGGER.Info("Server IP: " + CONF.SERVER.Addr)
	LOGGER.Info("Server Port: " + strconv.Itoa(CONF.SERVER.Port))
	fmt.Println(typeutils.Repeat("-", 70))
	LOGGER.Info("Loglevel: " + LOGGER.Level)
	LOGGER.Info("Log use file: " + strconv.FormatBool(LOGGER.UseFile))
	LOGGER.Info("Path to logfile: " + LOGGER.Path + "\\" + LOGGER.Filename)
	fmt.Println(typeutils.Repeat("-", 70))
}
