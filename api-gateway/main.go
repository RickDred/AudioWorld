package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	audio "proto/audio-proto"
	auth "proto/auth-proto"
)

type gatewayService struct{}

func main() {
	gateway := gatewayService{}
	port := ":50053"

	server := grpc.NewServer()
	RegisterGatewayServiceServer(server, &gateway)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("API Gateway started, listening on port", port)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *gatewayService) Download(ctx context.Context, req *audio.DownloadRequest) (*audio.DownloadResponse, error) {
	// Implement the logic to delegate the download request
	// to the Audio Service and return its response
	audioClient := audio.NewAudioServiceClient(audioConn) // Establish a gRPC connection to Audio Service
	res, err := audioClient.Download(ctx, req)
	if err != nil {
		log.Fatalf("Failed to call AudioService Download: %v", err)
	}

	return res, nil
}

func (s *gatewayService) Add(ctx context.Context, req *audio.AddRequest) (*audio.AddResponse, error) {
	// Implement the logic to delegate the add request
	// to the Audio Service and return its response
	audioClient := audio.NewAudioServiceClient(audioConn) // Establish a gRPC connection to Audio Service
	res, err := audioClient.Add(ctx, req)
	if err != nil {
		log.Fatalf("Failed to call AudioService Add: %v", err)
	}

	return res, nil
}

func (s *gatewayService) Listen(ctx context.Context, req *audio.ListenRequest) (*audio.ListenResponse, error) {
	// Implement the logic to delegate the listen request
	// to the Audio Service and return its response
	audioClient := audio.NewAudioServiceClient(audioConn) // Establish a gRPC connection to Audio Service
	res, err := audioClient.Listen(ctx, req)
	if err != nil {
		log.Fatalf("Failed to call AudioService Listen: %v", err)
	}

	return res, nil
}

func (s *gatewayService) Authorization(ctx context.Context, req *auth.AuthorizationRequest) (*auth.AuthorizationResponse, error) {
	// Implement the logic to delegate the authorization request
	// to the Auth Service and return its response
	authClient := auth.NewAuthServiceClient(authConn) // Establish a gRPC connection to Auth Service
	res, err := authClient.Authorization(ctx, req)
	if err != nil {
		log.Fatalf("Failed to call AuthService Authorization: %v", err)
	}

	return res, nil
}

func (s *gatewayService) Registration(ctx context.Context, req *auth.RegistrationRequest) (*auth.RegistrationResponse, error) {
	// Implement the logic to delegate the registration request
	// to the Auth Service and return its response
	authClient := auth.NewAuthServiceClient(authConn) // Establish a gRPC connection to Auth Service
	res, err := authClient.Registration(ctx, req)
	if err != nil {
		log.Fatalf("Failed to call AuthService Registration: %v", err)
	}

	return res, nil
}

//package main
//
//import (
//	"log"
//	"net"
//
//	"google.golang.org/grpc"
//)
//
//func failOnError(err error, msg string) {
//	if err != nil {
//		log.Panicf("%s: %s", msg, err)
//	}
//}
//
//// Server with embedded UnimplementedGreetServiceServer
//type Server struct {
//	greetpb.UnimplementedGreetServiceServer
//}
//
//func main() {
//	l, err := net.Listen("tcp", "0.0.0.0:50051")
//	failOnError(err, "Failed to listen")
//
//	s := grpc.NewServer()
//	greetpb.RegisterGreetServiceServer(s, &Server{})
//	log.Println("Server is running on port:50051")
//	if err := s.Serve(l); err != nil {
//		log.Fatalf("failed to serve:%v", err)
//	}
//}

//package main
//
//import (
//	pb "api-gateway/audio" // Import your gRPC generated code
//	"github.com/gin-gonic/gin"
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/credentials"
//	"log"
//	"net/http"
//)
//
//func main() {
//	// Create a new Gin router
//	router := gin.Default()
//
//	// Define your routes
//	router.Any("/users/*path", reverseProxy("localhost:8001", true))
//	router.Any("/products/*path", reverseProxy("localhost:8002", true))
//
//	// Start the server
//	log.Fatal(router.Run(":8080"))
//}
//
//func reverseProxy(targetHost string, enableTLS bool) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		// Create a connection to the gRPC backend service
//		var conn *grpc.ClientConn
//		var err error
//		if enableTLS {
//			creds, err := credentials.NewClientTLSFromFile("path/to/ssl/certificate.pem", "")
//			if err != nil {
//				log.Fatalf("Failed to create TLS credentials: %v", err)
//			}
//			conn, err = grpc.Dial(targetHost, grpc.WithTransportCredentials(creds))
//		} else {
//			conn, err = grpc.Dial(targetHost, grpc.WithInsecure())
//		}
//		if err != nil {
//			log.Fatalf("Failed to dial gRPC server: %v", err)
//		}
//		defer conn.Close()
//
//		// Create a reverse proxy with the gRPC connection
//		proxy := grpcHandlerFunc(conn)
//
//		// Serve the gRPC request through the reverse proxy
//		proxy.ServeHTTP(c.Writer, c.Request)
//	}
//}
//
//func grpcHandlerFunc(conn *grpc.ClientConn) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		// Create a gRPC client using the connection
//		client := pb.NewYourGRPCServiceClient(conn)
//
//		// Forward the request to the gRPC backend service
//		// You can make method calls on the gRPC client as needed
//		// Example:
//		// response, err := client.YourGRPCMethod(context.Background(), request)
//
//		// Handle the gRPC response and send it back as HTTP response
//		// Example:
//		// if err != nil {
//		//     http.Error(w, err.Error(), http.StatusInternalServerError)
//		//     return
//		// }
//		// w.Header().Set("Content-Type", "application/json")
//		// w.WriteHeader(http.StatusOK)
//		// json.NewEncoder(w).Encode(response)
//	}
//}

//package main
//
//import (
//	"context"
//	"log"
//	"net"
//
//	"google.golang.org/grpc"
//	"google.golang.org/grpc/codes"
//	"google.golang.org/grpc/status"
//
//	pb "api-gateway/audio" // Import the generated proto package
//
//	backendpb "api-gateway/backend" // Import the generated backend service package
//)
//
//const (
//	port         = ":50051"          // Specify the port on which the API gateway will listen
//	backendHost  = "localhost:50052" // Specify the backend service host and port
//	maxDataBytes = 1024 * 1024 * 5   // Maximum allowed audio file size in bytes (5 MB in this example)
//)
//
//type server struct {
//	backendClient backendpb.BackendServiceClient
//}
//
//func (s *server) Download(ctx context.Context, req *pb.DownloadRequest) (*pb.DownloadResponse, error) {
//	backendReq := &backendpb.DownloadRequest{
//		FileId: req.FileId,
//	}
//
//	backendRes, err := s.backendClient.Download(ctx, backendReq)
//	if err != nil {
//		return nil, status.Errorf(codes.Internal, "failed to download audio from backend: %v", err)
//	}
//
//	res := &pb.DownloadResponse{
//		AudioData: backendRes.AudioData,
//	}
//
//	return res, nil
//}
//
//func (s *server) Upload(ctx context.Context, req *pb.UploadRequest) (*pb.UploadResponse, error) {
//	if len(req.AudioData) > maxDataBytes {
//		return nil, status.Errorf(codes.InvalidArgument, "audio file size exceeds the limit")
//	}
//
//	backendReq := &backendpb.UploadRequest{
//		AudioData: req.AudioData,
//	}
//
//	backendRes, err := s.backendClient.Upload(ctx, backendReq)
//	if err != nil {
//		return nil, status.Errorf(codes.Internal, "failed to upload audio to backend: %v", err)
//	}
//
//	res := &pb.UploadResponse{
//		FileId: backendRes.FileId,
//	}
//
//	return res, nil
//}
//
//func main() {
//	// Set up gRPC server
//	listener, err := net.Listen("tcp", port)
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//
//	// Create a gRPC server instance
//	srv := grpc.NewServer()
//
//	// Create a backend service client
//	backendConn, err := grpc.Dial(backendHost, grpc.WithInsecure())
//	if err != nil {
//		log.Fatalf("failed to connect to backend service: %v", err)
//	}
//	backendClient := backendpb.NewBackendServiceClient(backendConn)
//
//	// Register the API gateway server
//	pb.RegisterAudioServiceServer(srv, &server{
//		backendClient: backendClient,
//	})
//
//	// Start the gRPC server
//	log.Printf("API gateway is listening on port %s", port)
//	if err := srv.Serve(listener); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}
//}
