from golang

COPY . .

RUN go build -v 

RUN go test -v