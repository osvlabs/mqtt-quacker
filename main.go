/*
 * mqtt-quacker
 *
 *
 * Contact: zhangwb@shinetechchina.com
 */

package main

import (
	"fmt"
	"mqtt-quacker/app"
	"os"
	"runtime"
)

func main() {
	fmt.Printf("Server started\n")

	mqttConfig := app.QuackerConfig{
		Host:     getEnv("QUACKER_HOST", "127.0.0.1"), // "mqtt.osvie.com",
		Port:     getEnv("QUACKER_PORT", "1883"),
		Username: getEnv("QUACKER_USERNAME", ""),
		Password: getEnv("QUACKER_PASSWORD", ""),
		Topic:    getEnvOrFail("QUACKER_TOPIC"),
		ClientId: getEnv("QUACKER_CLIENTID", "mqtt-quacker"),
		QoS:      getEnv("QUACKER_QOS", "0"),
		Interval: getEnv("QUACKER_INTERVAL", "1"),
		DataFile: getEnv("QUACKER_DATAFILE", "/data.json"),
		DryRun:   getEnv("QUACKER_DRYRUN", "") != "",
	}

	quacker := app.NewQuacker(mqttConfig)
	runtime.SetFinalizer(&quacker, func(obj *app.Quacker) {
		obj.Close()
	})

	quacker.Start()
	defer quacker.Close()
}

func getEnv(key string, defaultValue string) string {
	value, found := os.LookupEnv(key)
	if found == false {
		return defaultValue
	}
	return value
}

func getEnvOrFail(key string) string {
	value, found := os.LookupEnv(key)
	if found == false {
		panic("Need env key" + key)
	}
	return value
}
