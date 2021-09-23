FROM golang

WORKDIR /

COPY . .

RUN go mod download

EXPOSE 9000

CMD ["go", "run", "main.go"]