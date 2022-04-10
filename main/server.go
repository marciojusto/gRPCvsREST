package main

import (
	"encoding/json"
	"gRPCvsREST/grpc/pb"
	"gRPCvsREST/grpc/service"
	"gRPCvsREST/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
)

var courseList = model.NewCourses()

func main() {
	go startGrpc()
	http.HandleFunc("/course", CourseListHandler)
	http.ListenAndServe(":8889", nil)
}

func startGrpc() {
	lis, err := net.Listen("tcp", "localhost:50056")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	courseService := service.NewCourseGrpcService()
	courseService.Courses = courseList

	pb.RegisterCourseServiceServer(grpcServer, courseService)
	grpcServer.Serve(lis)
}

func CourseListHandler(w http.ResponseWriter, r *http.Request) {
	courseJson, _ := json.Marshal(courseList)
	w.Write([]byte(courseJson))
}
