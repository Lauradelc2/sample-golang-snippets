from golang

WORKDIR /src

COPY . .

RUN go get -d ./...

RUN go build -v ./...

RUN go test -v ./...