//package main
//
//import (
//	"context"
//	"fmt"
//	"log"
//	"net"
//	"strings"
//
//	"google.golang.org/grpc"
//)
//
//// Config represents the configuration for backend services
//type Config struct {
//	ServiceHosts map[string]string
//}
//
//// APIService defines the gRPC service
//type APIService struct {
//	config Config
//}
//
//// MyServiceServer is the interface that defines the methods of the backend service
//type MyServiceServer interface {
//	GetData(context.Context, *GetDataRequest) (*GetDataResponse, error)
//}
//
//// MyBackendService represents the backend service
//type MyBackendService struct{}
//
//// GetData handles the GetData gRPC method
//func (s *MyBackendService) GetData(ctx context.Context, req *GetDataRequest) (*GetDataResponse, error) {
//	// Handle the request and return the response
//	return &GetDataResponse{
//		Message: fmt.Sprintf("Hello, %s! This is the backend service.", req.Name),
//	}, nil
//}
//
//// NewAPIService creates a new instance of APIService
//func NewAPIService(config Config) *APIService {
//	return &APIService{config: config}
//}
//
//// ServeGRPC handles the incoming gRPC requests
//func (gw *APIService) ServeGRPC() {
//	lis, err := net.Listen("tcp", ":50051")
//	if err != nil {
//		log.Fatalf("Failed to listen: %v", err)
//	}
//
//	grpcServer := grpc.NewServer()
//	RegisterMyServiceServer(grpcServer, &MyBackendService{})
//
//	log.Println("gRPC server is listening on port 50051...")
//	if err := grpcServer.Serve(lis); err != nil {
//		log.Fatalf("Failed to serve: %v", err)
//	}
//}
//
//// ProxyGetData forwards the GetData gRPC request to the backend service
//func (gw *APIService) ProxyGetData(ctx context.Context, req *GetDataRequest) (*GetDataResponse, error) {
//	// Extract the service name from the request
//	parts := strings.SplitN(req.ServiceName, ".", 2)
//	if len(parts) < 2 {
//		return nil, fmt.Errorf("Invalid request: missing service name")
//	}
//	service := parts[0]
//
//	// Retrieve the backend service host from the config
//	host, ok := gw.config.ServiceHosts[service]
//	if !ok {
//		return nil, fmt.Errorf("Service not found: %s", service)
//	}
//
//	// Connect to the backend service
//	conn, err := grpc.Dial(host, grpc.WithInsecure())
//	if err != nil {
//		return nil, fmt.Errorf("Failed to connect to backend service: %v", err)
//	}
//	defer conn.Close()
//
//	// Create a gRPC client for the backend service
//	client := NewMyServiceClient(conn)
//
//	// Forward the request to the backend service
//	return client.GetData(ctx, req)
//}
//
//func main() {
//	// Define the configuration for backend services
//	config := Config{
//		ServiceHosts: map[string]string{
//			"service1": "localhost:8001",
//			"service2": "localhost:8002",
//			// Add more backend services here
//		},
//	}
//
//	// Create a new instance of APIService
//	apiService := NewAPIService(config)
//
//	// Start the gRPC server
//	go apiService.ServeGRPC()
//
//	// Create a gRPC client to make requests to the API Gateway
//	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
//	if err != nil {
//		log.Fatalf("Failed to connect to API Gateway: %v", err)
//	}
//	defer conn.Close()
//
//	client := NewMyServiceClient(conn)
//
//	// Make a request to the API Gateway, which will proxy the request to the backend service
//	req := &GetDataRequest{
//		ServiceName: "service1.GetData", // Specify the backend service and method
//		Name:        "John",
//	}
//
//	resp, err := client.ProxyGetData(context.Background(), req)
//	if err != nil {
//		log.Fatalf("Failed to get data from backend service: %v", err)
//	}
//
//	log.Printf("Response from backend service: %s", resp.Message)
//}

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// Config represents the configuration for backend services
type Config struct {
	ServiceHosts map[string]string
}

// APIGateway represents the API Gateway
type APIGateway struct {
	config Config
}

// NewAPIGateway creates a new instance of APIGateway
func NewAPIGateway(config Config) *APIGateway {
	return &APIGateway{config: config}
}

// ServeHTTP handles the incoming HTTP requests
func (gw *APIGateway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Extract the service name from the URL path
	parts := strings.SplitN(r.URL.Path, "/", 3)
	if len(parts) < 3 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	service := parts[1]

	// Retrieve the backend service host from the config
	host, ok := gw.config.ServiceHosts[service]
	if !ok {
		http.Error(w, "Service not found", http.StatusNotFound)
		return
	}

	// Reverse proxy the request to the appropriate backend service
	backendURL, _ := url.Parse(host)
	proxy := httputil.NewSingleHostReverseProxy(backendURL)
	proxy.ServeHTTP(w, r)
}

func main() {
	// Define the configuration for backend services
	config := Config{
		ServiceHosts: map[string]string{
			"service1": "http://localhost:8001",
			"service2": "http://localhost:8002",
			// Add more backend services here
		},
	}

	// Create a new instance of APIGateway
	gateway := NewAPIGateway(config)

	// Start the API Gateway server
	port := 8080
	fmt.Printf("API Gateway is listening on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), gateway))
}
