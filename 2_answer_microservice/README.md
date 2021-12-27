# Microservice REST API & gRPC | Golang


Microservice with gRPC & REST API writen golang with:
- Docker environment with compose integrate 
- MYSQL
- GO GIN Framework (REST)
- GO gRPC proto (gRPC)
- Unit Testing
- DDD (Domain Driven Design) approach

Arsitecture:


# Getting start

Clone repository

```
git clone
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