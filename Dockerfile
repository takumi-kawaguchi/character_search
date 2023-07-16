FROM golang:1.20

WORKDIR /app

# COPY ./go.mod /app/
# RUN go mod tidy

CMD [ "go", "run", "main.go" ]