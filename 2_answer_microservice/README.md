# Microservice REST API & gRPC | Golang


Microservice with gRPC & REST API writen golang with:
- Docker environment with compose integrate 
- MYSQL
- GO GIN Framework (REST)
- GO gRPC proto (gRPC)
- Unit Testing
- DDD (Domain Driven Design) approach

Arsitecture:
![Arsitecture](https://raw.githubusercontent.com/undercode99/test-golang-submision/main/2_answer_microservice/flowmicro.png)




# Getting start

Clone repository

```
git clone https://github.com/undercode99/test-golang-submision.git
cd 2_answer_microservice
```

Build microservices
``` 
docker-compose up --build
```

Running up Microservices
``` 
docker-compose up -d 
```

Running unit test & integration test
``` 
docker-compose -f docker-compose.test.yaml up --build --abort-on-container-exit 
```
