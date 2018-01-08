FROM golang:1.7.3 as builder
LABEL maintainer="Solomon White <rubysolo@gmail.com>"

WORKDIR /src
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o spa_envy .

FROM scratch
WORKDIR /
COPY --from=builder /src/spa_envy .

EXPOSE 3000

CMD ["./spa_envy"]  
