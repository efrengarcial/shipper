package main

import (
	"context"
	"log"

	pb "github.com/efrengarcial/shipper/user-service/proto/auth"
	"github.com/micro/go-micro"
	"time"
)

const topic = "user.created"

type Subscriber struct{}

func (sub *Subscriber) Process(ctx context.Context, user *pb.User) error {
	log.Println("Picked up a new message")
	log.Println("Sending email to:", user.Name)
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.email"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	srv.Init()

	micro.RegisterSubscriber(topic, srv.Server(), new(Subscriber))

	// Run the server
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
