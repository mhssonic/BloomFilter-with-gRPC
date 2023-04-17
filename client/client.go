package main

import (
	"fmt"
	"log"

	"BloomFilter/request"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	added := []string{"hey", "ali", "pashmaki", "sfds", "sdidd", "reza", "hi", "dude", "nun", "tof"}
	notAdded := []string{"hy", "hadi", "shokolat", "wqeds", "hgfde", "hasan", "james", "uwuwu", "12312f", "12edw"}

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := request.NewBloomFilterClient(conn)

	for _, str := range added {
		_, err := c.AddString(context.Background(), &request.Message{Body: str})
		if err != nil {
			log.Fatalf("Error when calling AddString: %s", err)
		}
	}
	for _, str := range added {
		response, err := c.IsThere(context.Background(), &request.Message{Body: str})
		if err != nil {
			log.Fatalf("Error when calling IsThere: %s", err)
		}
		fmt.Printf("%s %t\n", str, response.Body)
	}
	fmt.Printf("-------------------\n")
	for _, str := range notAdded {
		response, err := c.IsThere(context.Background(), &request.Message{Body: str})
		if err != nil {
			log.Fatalf("Error when calling IsThere: %s", err)
		}
		fmt.Printf("%s %t\n", str, response.Body)
	}
}
