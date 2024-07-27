FROM golang:alpine AS builder

WORKDIR /time-tracker

ADD ../go.mod .

COPY . .

RUN go build -o ./build/time.exe ./cmd/main.go

FROM alpine

WORKDIR /time-tracker

COPY ./configs/.env /time-tracker/configs/.env
COPY --from=builder /time-tracker/build /time-tracker

CMD ["/time-tracker/time.exe"]