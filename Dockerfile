FROM golang:1.18

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY .env /app
RUN go mod download

#COPY *.go ./
COPY . ./
EXPOSE 8891

RUN go build -o /ews

CMD [ "/ews" ]