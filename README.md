# Blogger

This is just a basic CRUD API written in Golang, implemented using gRPC and MongoDB.

A blog can be created, read, updated, and deleted with this API. Finally, you can also view (scan for) a list of existing blogs in your databases.

## How To Use
This assumes you have Golang and MongoDB installed on your local machine.

The easiest way to try out this gRPC API is to use an existing gRPC client, such as Evans CLI, which can be found at https://github.com/ktr0731/evans

From there, you will be able to inspect the available packages, services, and RPCs within this project, notably:
* Create a blog
* Read a blog
* Update a blog
* Delete a blog
* Scan a list of blogs
-----------------------
1. Run `generate.sh`
2. Make sure your local MongoDB is running.
3. Utilize the Evans CLI

## Goals of this Project
Of course, this is just a basic project, intended to get my toes wet in the world of gRPC. I wanted to check out the differences between gRPC and REST API, and see why Google is leveraging the former.

From this project, I was able to:
* Learn the gRPC theory
* Compare gRPC and REST API paradigms
* Get more practice developing in Golang
* Get a hands-on introduction to Protocol Buffers
* Implement Unary, Server Streaming, Client Streaming, and Bi-Directional Streaming APIs and learn the different use cases for each.
* Get more practice with Golang contexts

## Future Plans
Much more can be done with this project, and are coming in the future (to be updated):
* Front End client with friendly UI for usability

## Sources
This project was creating while following the below course by Stephane Maarek:
https://www.udemy.com/course/grpc-golang/