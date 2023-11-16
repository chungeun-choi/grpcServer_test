package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

/*
 grpc 구현 순서 https://grpc.io/docs/languages/go/basics/
	1) protobuf 파일 정의
		1-1) 주고 받을 데이터의 규격 정의
		1-2) 어떠한 메소드를 통해 1-1 항목에서 정의한 데이터를 주고받을 것인지 정의
	2) protoc 명령어를 통해 pb 파일 인코딩?(인코딩이라는 표현이 맞나?)
	3) 생성된 pb.go 파일들의 실제 구현체를 App프로젝트 내에서 구현이 필요
*/
import (
	"context"
	pb "grpcApp/proto"
)

// proto 파일에 명시되어진 interface를 구현하기위한 구조체를 사용해야함
type ChatServer struct {
	pb.UnimplementedChatServiceServer
}

/*
해당 구조체를 가지고있는 추상 함수를 사용하는 서버에서 실제 구현이 이루어져야함
실제 비즈니스 로직이 담기는 부분이라 판단됨 (흔히 서비스라고 구현하는 객체)
*/
func (cs *ChatServer) SayHello(ctx context.Context, message *pb.Message) (*pb.Message, error) {
	log.Printf("Received: %v", message.Body)

	result := &pb.Message{
		Body: "Recevied s",
	}
	return result, nil
}

/*
grpc 서버 구현 listen 형태
 1. tcp 통신을 위한 네트워크 연결 객체 생성
 2. grpc 섭시를 실제로 담당하게 될 구조체를 객체로 초기화
 3. grpc Server 객체를 생성
 4. proto 파일을 통해 generate 한 pb.go 파일에서 Register 함수를 통해
    2번 항목과 3번항목을 입력을 통해 의존성 주입
 5. grpc Server의 Serve 함수를 통해 앞서 1번 항목에서 만든 네트워크 객체 의존성 주입
*/
func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}

	chatSever := ChatServer{}
	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, &chatSever)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
