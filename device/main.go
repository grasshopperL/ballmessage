package main

import (
	"github.com/shirou/gopsutil/v3/process"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"path"
	"sync"
	"syscall"
)

// getConfig get config value from config file according key
func getConfig(key string) (bool, interface{}) {

	defer configLock.Unlock()
	configLock.Lock()
	return viper.IsSet(key), viper.Get(key)
}

// setConfig write config value to config file according key
func setConfig(key string, value interface{}) error {

	defer configLock.Unlock()
	configLock.Lock()
	viper.Set(key, value)
	return viper.WriteConfigAs(path.Join(configDir, configName))
}

// findProcess finds pid according to process name
func findProcess(name string) int {

	processes, err := process.Processes()
	if err != nil {
		return -1
	}
	for _, p := range processes {
		if n, err := p.Name(); err == nil && n == name {
			return int(p.Pid)
		}
	}
	return -1
}

// killProcess kill process according to process name
func killProcess(name string) error {

	processes, err := process.Processes()
	if err != nil {
		return err
	}
	for _, p := range processes {
		if n, err := p.Name(); err == nil && n == name {
			return p.Kill()
		}
	}
	return nil
}

// configLock is mutex lock for config file
var configLock sync.Mutex
var configDir string
var configName string

func main() {

	log.Println("<INFO> [ main ] = {start}")
	chanOS := make(chan os.Signal, 1)
	signal.Notify(chanOS, os.Interrupt, syscall.SIGTERM)
	e := <-chanOS
	log.Println("<INFO> [main] = {exit}", e)
}
