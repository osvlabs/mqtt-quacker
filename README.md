# mqtt-quacker

Send mock data to your MQTT topic.

## Usage docker
`
docker run -e QUACKER_HOST=mqtt.host.com -e QUACKER_PORT=1883 -e QUACKER_USERNAME=mqtt-username -e QUACKER_PASSWORD=mqtt-password -e QUACKER_TOPIC=my-topic/telemetry -v /home/zgldh/my-project/data.json:/data.json zgldh/mqtt-quacker
`

## Usage docker-compose

Edit the docker-compose.yml  
```
docker-compose up 
```


## Variables

name| descrpition | sample
----|-------------|---------
QUACKER_HOST| The host to your MQTT server. | "mqtt.host.com"
QUACKER_PORT| The MQTT server port. |"1883"
QUACKER_USERNAME| For MQTT server auth. |"mqtt-username"
QUACKER_PASSWORD| For MQTT server auth. |"mqtt-password"
QUACKER_TOPIC| Which topic do you want the mock data send to? |"your/topic/to/send"
QUACKER_CLIENTID| The client ID |"mqtt-quacker"
QUACKER_QOS| Please check MQTT doc. 0, 1, 2 |"0"
QUACKER_INTERVAL| Time interval between two data sending. |"1"
QUACKER_DATAFILE| The mock data template. |"/data.json"