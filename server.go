package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"
	"sync"

	pb "github.com/ssummers02/go_grpc"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func FindByPattern(pattern string, body string) string {
	re := regexp.MustCompile(pattern)
	sub := re.FindAllStringSubmatch(body, -1)
	for _, element := range sub {
		return element[1]
	}

	return ""
}

const URL = "https://www.rusprofile.ru/search?query="

type Greeter struct {
	wg sync.WaitGroup
}

func (g *Greeter) FirmInfoGet(ctx context.Context, request *pb.FirmByINNRequest) (*pb.FirmInfoResponse, error) {
	response := &pb.FirmInfoResponse{
		Name: "",
		Inn:  "",
		Kpp:  "",
		Boss: ""}
	resp, err := http.Get(URL + request.GetInn())
	if err != nil {
		return response, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	var bodyToParse = string(body)

	response.Inn = FindByPattern(".*clip_kpp.?>(.*)<", bodyToParse)
	response.Kpp = FindByPattern(".*clip_kpp.?>(.*)<", bodyToParse)

	var result = FindByPattern("legalName.?>(.*)</", bodyToParse)
	if result != "" {
		response.Name = strings.Replace(result, "&quot;", "\"", -1)
		response.Boss = FindByPattern(fmt.Sprintf("<meta name=\"keywords\" content=.*%v, (.*), ИНН ", result), bodyToParse)
	}

	return response, nil
}

// New creates new server greeter
func New() *Greeter {
	return &Greeter{}
}

// Start starts server
func (g *Greeter) Start() {
	g.wg.Add(1)
	go func() {
		log.Fatal(g.startGRPC())
		g.wg.Done()
	}()
	g.wg.Add(1)
	go func() {
		log.Fatal(g.startREST())
		g.wg.Done()
	}()
}
func (g *Greeter) WaitStop() {
	g.wg.Wait()
}
func (g *Greeter) startGRPC() error {
	lis, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, g)
	grpcServer.Serve(lis)
	return nil
}
func (g *Greeter) startREST() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterGreeterHandlerFromEndpoint(ctx, mux, ":8090", opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}
