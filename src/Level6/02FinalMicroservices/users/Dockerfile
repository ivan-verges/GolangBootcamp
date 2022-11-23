FROM golang:1.19.3-alpine

RUN mkdir -p /usr/local/go/src/BairesDev/Level6/02_Final_Microservices/users/data
RUN mkdir -p /usr/local/go/src/BairesDev/Level6/02_Final_Microservices/users/grpc
RUN mkdir -p /usr/local/go/src/BairesDev/Level6/02_Final_Microservices/users/main
RUN mkdir -p /usr/local/go/src/BairesDev/Level6/02_Final_Microservices/users/model
RUN mkdir -p /usr/local/go/src/BairesDev/Level6/02_Final_Microservices/users/rest

WORKDIR /usr/local/go/src/BairesDev/Level6/02_Final_Microservices/users

COPY users/data data/
COPY users/grpc grpc/
COPY users/main main/
COPY users/model model/
COPY users/rest rest/

RUN go mod init users
RUN go mod tidy
RUN go mod download

WORKDIR /usr/local/go/src/BairesDev/Level6/02_Final_Microservices/users/main

RUN go build -o /users

EXPOSE 8001
EXPOSE 9001

CMD [ "/users" ]
