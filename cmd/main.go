package main

import (
	"context"
	"flag"
	"log"
	"time"

	v1 "github.com/nickstern2002/gRPCtestingServer/pkg/protogen/compute/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:8080", "http service address")
	mood = flag.Int("mood", 0, "mood level")
)

func main() {
	flag.Parse()

	for {
		conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Println("Failed to create client:", err)
			time.Sleep(5 * time.Second)
			continue
		}
		defer conn.Close()

		client := v1.NewJunkyardServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		// Server Response
		resp, err := client.MakeMyDayBetter(ctx, &v1.MakeMyDayBetterRequest{Mood: int32(*mood)})
		if err != nil {
			log.Println("Error calling MakeMyDayBetter:", err)
		} else {
			log.Println(resp)
		}

		time.Sleep(5 * time.Second)
	}
}
