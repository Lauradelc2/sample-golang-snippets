from golang

COPY . .

RUN go get 

RUN go build -v 

RUN go test -v