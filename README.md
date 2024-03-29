# mqtt-quacker

Send mock data to your MQTT topic.

## Usage docker
```
docker run 
  -e QUACKER_HOST=mqtt.host.com 
  -e QUACKER_PORT=1883 
  -e QUACKER_USERNAME=mqtt-username 
  -e QUACKER_PASSWORD=mqtt-password 
  -e QUACKER_TOPIC=my-topic/telemetry 
  -v /home/zgldh/my-project/data.json:/data.json 
  zgldh/mqtt-quacker:1.2
```

## Usage docker-compose

Edit the docker-compose.yml  
```
docker-compose up 
```


## Variables

name| descrpition | sample
----|-------------|---------
QUACKER_HOST| The host to your MQTT server. | `mqtt.host.com`
QUACKER_PORT| The MQTT server port. |`1883`
QUACKER_USERNAME| For MQTT server auth. |`mqtt-username`
QUACKER_PASSWORD| For MQTT server auth. |`mqtt-password`
QUACKER_TOPIC| Which topic do you want the mock data send to? |`your/topic/to/send`
QUACKER_CLIENTID| The client ID |`mqtt-quacker`
QUACKER_QOS| Please check MQTT doc. 0, 1, 2 |`0`
QUACKER_INTERVAL| Time interval between two data sending. (in ms) |`1000`
QUACKER_DATAFILE| The mock data template. |`./data.json`
QUACKER_DRYRUN| Dont push to server, just output payload. |""

## Custom Data
Please edit the file `data.json` to any text you want. It supports following placeholders:
- `q:float:{min},{max}` to generate a float number between [min, max).
- `q:int:{min},{max}` to generate an integer number between [min, max).
- `q:string:{str1},{str2},{str3},...,{strN}` to get one string from n strings randomly.
- `q:timestamp` to get a current timestamp like `1620904435`.
- `q:datetime` to get a current datetime string. You can custom the date format: `q:datetime:Mon Jan _2 15:04:05 MST 2006` Please refer to: https://golang.org/src/time/format.go

Currently, no more placeholders supported.

## Update Logs
1.2
- Breaking changes:  
    Use `{q:datetime}` instead of `{q:timestamp}` 
    The `{q:timestamp}` is changed to output real timestamp long integer like `1620904435`
    
1.1
- Added `{q:timestamp}` token.

1.0
- The beginning.
