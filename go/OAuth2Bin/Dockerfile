FROM golang:1.13
LABEL maintainer=""

WORKDIR /app
COPY . .

RUN go build -o oauth2bin .
CMD [ "./oauth2bin" ]
