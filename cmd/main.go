package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/atye/wikitable-api/service"
	"github.com/atye/wikitable-api/service/pb"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func main() {
	var eg errgroup.Group

	lis, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}

	svc := &service.Service{}
	svr := grpc.NewServer()

	pb.RegisterWikiTableServer(svr, svc)
	eg.Go(func() error {
		return svr.Serve(lis)
	})

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	rMux := runtime.NewServeMux()

	err = pb.RegisterWikiTableHandlerFromEndpoint(context.Background(), rMux, "127.0.0.1:2000", opts)
	if err != nil {
		log.Fatal(err)
	}

	mux := setupHTTPMux(rMux)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatal("PORT env not set")
	}

	eg.Go(func() error {
		return http.ListenAndServe(":"+port, mux)
	})

	log.Fatal(eg.Wait())
}

func setupHTTPMux(rMux *runtime.ServeMux) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./static"))))
	mux.Handle("/api/", rMux)
	mux.HandleFunc("/swagger.json", serveSwagger)
	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui", http.FileServer(http.Dir("swagger/swagger-ui"))))

	return mux
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "swagger/apidocs.swagger.json")
}
