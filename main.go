package main

import messenger "github.com/AbdouTlili/go-messenger"

func main() {

	user := messenger.User{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

}
