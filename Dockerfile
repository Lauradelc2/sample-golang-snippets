# Este comentario será inserido apenas como marcador
# para um novo pull-request, solicitado pelo exercicio.

from golang as build

WORKDIR /src

COPY greet/ .

RUN go get .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/greet

RUN go test -v ./...

from alpine

COPY --from=build /bin/greet /bin/greet

COPY --from=build /src/etc/greet-api.yaml /etc/greet-api.yaml

ENTRYPOINT ["/bin/greet", "-f", "/etc/greet-api.yaml"]