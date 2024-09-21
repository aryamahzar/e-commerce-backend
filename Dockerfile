FROM golang:1.22 AS build

# Set to your project's root directory
WORKDIR /go/src/e-commerce-backend 

# Copy your module files first
COPY go.mod go.sum ./ 

# Download dependencies based on go.mod
RUN go mod download 

COPY go ./go
COPY main.go .

ENV CGO_ENABLED=0

RUN go build -a -installsuffix cgo -o swagger .

FROM scratch AS runtime
COPY --from=build /go/src/e-commerce-backend/swagger ./
EXPOSE 8080/tcp
ENTRYPOINT ["./swagger"]