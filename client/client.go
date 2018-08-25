package main
  
//client.go
  
import (
    "log"
    "os"
      "time"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    pb "GRPC/helloworld"
)
  
const (
    address     = "localhost:50051"
    defaultName = "client  Requst"
)
  
func main() {
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Fatal("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewGreeterClient(conn)
  
    name := defaultName
    if len(os.Args) >1 {
        name = os.Args[1]
    }
    for true{
    r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
    if err != nil {
        log.Fatal("could not greet: %v", err)
    }
    log.Printf("Greeting: %s", r.Message)
         time.Sleep(2*time.Second)
   }
}