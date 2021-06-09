package main

import (
	"context"
	"fmt"
	"io"
	"justinschen/grpc-blog-crud/blog/blogpb"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Printf("%v - Starting Blog Client\n", time.Now())

	options := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", options)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	client := blogpb.NewBlogServiceClient(cc)

	fmt.Printf("%v - Creating Blog\n", time.Now())

	blog := &blogpb.Blog{
		AuthorId: "Justin",
		Title: "My First Blog",
		Content: "Content of First Blog",
	}

	// CREATE
	createBlogResp, err := client.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}
	blogID := createBlogResp.GetBlog().GetId()

	fmt.Printf("Blog has been created: %v\n", createBlogResp)

	// READ
	_, err2 := client.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "test"})
	if err2 != nil {
		fmt.Printf("Error during ReadBlog: %v\n", err2)
	}

	readBlogReq := &blogpb.ReadBlogRequest{BlogId: createBlogResp.GetBlog().Id}
	readBlogResp, err3 := client.ReadBlog(context.Background(), readBlogReq)
	if err3 != nil {
		fmt.Printf("Error during ReadBlog: %v\n", err2)
	}

	fmt.Printf("Blog was read: %v\n", readBlogResp)

	// UPDATE
	newBlog := &blogpb.Blog {
		Id: blogID,
		AuthorId: "Mugen Hanzo",
		Title: "Mugen Si",
		Content: "Honda Civic 2008 Mugen Si",
	}
	updateResp, updateErr := client.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{
		Blog: newBlog,
	})
	if updateErr != nil {
		fmt.Printf("Error occurred while updating blog: %v\n", updateErr)
	}
	fmt.Printf("Blog was updated: %v\n", updateResp)

	// DELETE
	deleteResp, deleteErr := client.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{
		BlogId: blogID,
	})
	if deleteErr != nil {
		fmt.Printf("Error occurred while deleting blog: %v\n", deleteErr)
	}
	fmt.Printf("Blog was deleted: %v\n", deleteResp)

	// LIST
	stream, err := client.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if err != nil {
		log.Fatalf("Error while calling ListBlog: %v\n", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while receiving from stream: %v\n", err)
		}
		fmt.Println(res.GetBlog())
	}
}