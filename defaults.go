package micro

import (
	"github.com/geiqin/go-micro/client"
	"github.com/geiqin/go-micro/debug/trace"
	"github.com/geiqin/go-micro/server"
	"github.com/geiqin/go-micro/store"

	// set defaults
	gcli "github.com/geiqin/go-micro/client/grpc"
	memTrace "github.com/geiqin/go-micro/debug/trace/memory"
	gsrv "github.com/geiqin/go-micro/server/grpc"
	memoryStore "github.com/geiqin/go-micro/store/memory"
)

func init() {
	// default client
	client.DefaultClient = gcli.NewClient()
	// default server
	server.DefaultServer = gsrv.NewServer()
	// default store
	store.DefaultStore = memoryStore.NewStore()
	// set default trace
	trace.DefaultTracer = memTrace.NewTracer()
}
