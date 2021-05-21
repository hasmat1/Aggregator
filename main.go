package main

import (
	"GoVaccineUpdaterNotifier/Service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "50001"
		log.Println("PORT not set")
	}
	log.Println("Using port", port)
	host := ""
	lis, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	defer func() {
		err = lis.Close()
		if err != nil {
			log.Fatalln("Failed to close listener:", err)
		}
	}()
	gRPCServer := grpc.NewServer()
	server := Service.NewServer()
	Service.RegisterEndpointsServer(gRPCServer, &server)
	go func() {
		if err := gRPCServer.Serve(lis); err != nil {
			log.Println("Failed to serve:", err)
		}
	}()
	log.Printf("Server successfully started on port: %s\n\n", port)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Printf("\nStopping server...\n")
	gRPCServer.Stop()
	log.Printf("Stopped server successfully")
}
