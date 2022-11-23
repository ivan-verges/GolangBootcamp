FROM golang:1.19.3-alpine

RUN mkdir -p /usr/local/go/src/BairesDev/Level6/02_Final_Microservices/questions/data
RUN mkdir -p /usr/local/go/src/BairesDev/Level6/02_Final_Microservices/questions/grpc
RUN mkdir -p /usr/local/go/src/BairesDev/Level6/02_Final_Microservices/questions/main
RUN mkdir -p /usr/local/go/src/BairesDev/Level6/02_Final_Microservices/questions/model
RUN mkdir -p /usr/local/go/src/BairesDev/Level6/02_Final_Microservices/questions/rest

WORKDIR /usr/local/go/src/BairesDev/Level6/02_Final_Microservices/questions

COPY questions/data data/
COPY questions/grpc grpc/
COPY questions/main main/
COPY questions/model model/
COPY questions/rest rest/

RUN go mod init questions
RUN go mod tidy
RUN go mod download

WORKDIR /usr/local/go/src/BairesDev/Level6/02_Final_Microservices/questions/main

RUN go build -o /questions

EXPOSE 8002
EXPOSE 9002

CMD [ "/questions" ]
