package main

import (
	"fmt"

	pb "example.com/tutorialpb"
)

func main() {
	test()
	fmt.Println("Starting to use protobuffer");
	p := pb.Person {
		Id: 1234,
		Name: "Sudeep Patel",
		Email: "test@example.com",
		Phones: []*pb.Person_PhoneNumber{
		{Number:"555-321", Type: pb.Person_HOME},

		},
	}
	var x = [2]int{1, 2};
	var y = [3]int{1, 2, 3};

	fmt.Println(x)
	// fmt.Println(x == y); //Not allow mismatch

	
	// [START marshal_proto]
	book1 := &pb.AddressBook{
		
	};

	var a  [10]pb.Person;
	fmt.Println(a)

	a[0] = p;
	// fmt.Println(a)
	fmt.Println(book1);

	// book1.People = *p;
	// out, err := proto.Marshal(book1);

	// if err != nil {
	// 	log.Fatalln("Failed to encode address book:", err)

	// }
	
}

func stringp(s pb.Person) *pb.Person {
	return &s
}