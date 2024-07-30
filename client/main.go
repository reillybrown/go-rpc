// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/reillybrown/go-rpc/contracts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	msg  = flag.String("msg", "", "a message to send")
)

func main() {
	// parse all command line args
	flag.Parse()
	// connect to server
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMessengerClient(conn)

	// contact server and print response
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if *msg != "" {
		_, err := c.AddMessage(ctx, &pb.Message{
			Body: *msg,
			Ts:   timestamppb.Now(), // example date handling
		})
		if err != nil {
			log.Fatalf("failed to send message: %v", err)
		}
	}
	r, err := c.ListMessages(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("failed to list messages: %v", err)
	}
	log.Printf("Messages %s", r.GetMessages())
}
