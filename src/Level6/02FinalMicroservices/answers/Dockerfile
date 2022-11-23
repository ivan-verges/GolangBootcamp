FROM golang:1.19.3-alpine

RUN mkdir -p /usr/local/go/src/BairesDev/Level6/02_Final_Microservices/answers/data
RUN mkdir -p /usr/local/go/src/BairesDev/Level6/02_Final_Microservices/answers/grpc
RUN mkdir -p /usr/local/go/src/BairesDev/Level6/02_Final_Microservices/answers/main
RUN mkdir -p /usr/local/go/src/BairesDev/Level6/02_Final_Microservices/answers/model
RUN mkdir -p /usr/local/go/src/BairesDev/Level6/02_Final_Microservices/answers/rest

WORKDIR /usr/local/go/src/BairesDev/Level6/02_Final_Microservices/answers

COPY answers/data data/
COPY answers/grpc grpc/
COPY answers/main main/
COPY answers/model model/
COPY answers/rest rest/

RUN go mod init answers
RUN go mod tidy
RUN go mod download

WORKDIR /usr/local/go/src/BairesDev/Level6/02_Final_Microservices/answers/main

RUN go build -o /answers

EXPOSE 8003
EXPOSE 9003

CMD [ "/answers" ]
