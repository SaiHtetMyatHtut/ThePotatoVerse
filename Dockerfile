FROM public.ecr.aws/docker/library/golang:latest

WORKDIR /app

COPY . ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /grumman_bot

CMD ["/grumman_bot"]