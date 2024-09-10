from golang

WORKDIR /src

COPY greet/ greet

RUN go get .

RUN go build -v ./...

RUN go test -v ./...