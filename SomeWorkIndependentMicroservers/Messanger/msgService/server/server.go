package main

import (
	"context"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "Messanger/messenger"
)

const (
	rabbitMQURL = "amqp://guest:guest@localhost:5672/"
	queueName   = "messages"
)

type server struct{}

func (s *server) SendMessage(ctx context.Context, req *pb.Message) (*pb.SendMessageResponse, error) {
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(req.Body),
		},
	)
	if err != nil {
		return nil, err
	}

	return &pb.SendMessageResponse{
		Success: true,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMessengerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
