FROM golang:alpine AS build
WORKDIR /app
COPY . .
RUN go build -o courtside-api .

FROM alpine
WORKDIR /app
COPY --from=build /app/courtside-api .
EXPOSE 8080
CMD ["./courtside-api"]
