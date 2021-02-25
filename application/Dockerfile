FROM hyperledger/fabric-tools:1.4.3

WORKDIR /srv

COPY . .

RUN go build main.go

CMD ["./main"]
