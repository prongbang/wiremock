# make push tag=1.0.0
push:
	docker build -t prongbang/wiremock:$(tag) .
	docker tag prongbang/wiremock:$(tag) prongbang/wiremock:$(tag)
	docker push prongbang/wiremock:$(tag)

# make login
login:
	curl -X POST -H 'Api-Key: ed2b7d14-3999-408e-9bb8-4ea739f2bcb5' -d '{"username":"admin", "password":"pass"}'  http://localhost:8000/api/v1/login

# make load_test
load_test:
	wrk -c 10000 -d 60 -t 4 http://localhost:8000/