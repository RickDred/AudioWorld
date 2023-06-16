FROM golang:latest
COPY ./ ./
RUN go build  -o user_client .
CMD ["./user_client"]