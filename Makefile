# make push tag=1.0.0
push:
	docker build -t prongbang/wiremock:$(tag) .
	docker tag prongbang/wiremock:$(tag) prongbang/wiremock:$(tag)
	docker push prongbang/wiremock:$(tag)