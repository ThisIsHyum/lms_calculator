FROM golang:alpine

WORKDIR /lms-calculator

COPY . .

RUN go build ./cmd/agent
RUN go build ./cmd/orchestrator