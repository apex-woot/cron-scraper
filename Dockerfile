FROM golang:1.22-alpine
WORKDIR /app
COPY ./scraper/go.mod .
CMD go mod download
COPY . .
RUN go build -o ./out/dist ./scraper/cmd/main.go
CMD ./out/dist
