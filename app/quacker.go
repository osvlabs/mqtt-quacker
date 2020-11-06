package app

import (
	"fmt"
	"math"
	"strconv"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// QuackerConfig - Configuration of MQTT server
type QuackerConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Topic    string
	ClientId string
	QoS      string
	Interval string // Interval - Seconds between two publish
	DataFile string // DataFile - Data template file path
	DryRun   bool
}

// Quacker - The quacker class.
type Quacker struct {
	config  QuackerConfig
	builder DataBuilder
}

// NewQuacker - Create a new Quacker object
func NewQuacker(config QuackerConfig) Quacker {
	return Quacker{
		config:  config,
		builder: NewDataBuilder(DataBuilderConfig{Path: config.DataFile}),
	}
}

// Close - Close the quacker mission
func (q *Quacker) Close() {
}

// Start - Start to subscribe MQTT and tranfer data into Pgsql
func (q *Quacker) Start() error {
	fmt.Printf("Quacker starting...\n")

	qos, err := strconv.Atoi(q.config.QoS)
	if err != nil {
		qos = 0
	}
	fmt.Printf("QoS %d\n", qos)

	interval, err := strconv.Atoi(q.config.Interval)
	if err != nil {
		interval = 1
	}
	interval = int(math.Max(float64(interval), 1))

	payload := ""

	opts := MQTT.NewClientOptions()
	opts.AddBroker(q.config.Host + ":" + q.config.Port)
	opts.SetClientID(q.config.ClientId)
	opts.SetUsername(q.config.Username)
	opts.SetPassword(q.config.Password)

	client := MQTT.NewClient(opts)
	publishLabel := "Publish"
	if q.config.DryRun {
		fmt.Println("Dry run")
		publishLabel = "Dry Run"
	} else {
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}

	fmt.Printf("MQTT server %s:%s\n", q.config.Host, q.config.Port)
	fmt.Printf("Client ID %s\n", q.config.ClientId)
	fmt.Println("Publisher Started to: " + q.config.Topic)
	for true {
		fmt.Printf("%s ---- %s ----\n", time.Now(), publishLabel)
		payload = q.getPayload()

		if !q.config.DryRun {
			token := client.Publish(q.config.Topic, byte(qos), false, payload)
			token.Wait()
		}

		fmt.Println(payload)
		time.Sleep(time.Second * time.Duration(interval))
	}

	if q.config.DryRun {
		fmt.Println("Publisher Terminated")
	} else {
		client.Disconnect(250)
		fmt.Println("Publisher Disconnected")
	}

	return nil
}

// getPayload - Get payload.
func (q *Quacker) getPayload() string {
	payload, err := q.builder.Make()
	if err != nil {
		panic(err)
	}
	return payload
}
