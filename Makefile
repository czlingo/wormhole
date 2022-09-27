build-docker:
	go mod vendor
	sudo docker build -t test .
	rm -rf vendor
	sudo docker images | grep 'none' | awk '{print $$3}' | xargs sudo docker image rm