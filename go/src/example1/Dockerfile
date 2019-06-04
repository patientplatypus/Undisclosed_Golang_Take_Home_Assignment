

FROM golang:latest as builder
COPY . /go/src/github.com/patientplatypus/project
WORKDIR /go/src/github.com/patientplatypus/project
RUN go get github.com/gorilla/handlers
RUN go get -u github.com/gorilla/mux
RUN go get github.com/rs/cors
RUN go test github.com/patientplatypus/project/controller/
RUN CGO_ENABLED=0 go build main.go

FROM alpine:3.7
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/patientplatypus/project/main .
EXPOSE 8000
CMD ["./main"]