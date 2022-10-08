image_repo = registry.cn-chengdu.aliyuncs.com/czlingo
image_name = wormhole
image_version = test

build-docker:
	go mod vendor
	sudo docker build -t $(image_repo)/$(image_name):$(image_version) .
	rm -rf vendor
	sudo docker images | grep 'none' | awk '{print $$3}' | xargs sudo docker image rm