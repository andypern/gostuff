# Hook to kafka

After you get submodules, etc:

	export GOPATH=~/gostuff/vendor

	cd ~/gostuff/hook-to-kafka

	go build
	

usage:

```
./hook-to-kafka:
 -kafkaHost string
	   space separated host:port pairs for kafka brokers (default "localhost:9092 localhost:9093")
 -listenPort string
	   port to listen on (default "8087")
 -topic string
	   topic you want to produce to (default "mytopic")

   ```
