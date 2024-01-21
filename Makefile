all:
	go build

show: all
	./note show

show-tag: all
	./note show -t general

add: all
	./note add "test4"

add-tag: all
	./note add "cp" -t cmd

remove: all
	./note remove 3

clean:
	rm ./note