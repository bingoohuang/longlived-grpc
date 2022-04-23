package resolver

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"google.golang.org/grpc/resolver"
)

func init() {
	resolver.Register(&StaticBuilder{})
}

type StaticBuilder struct{}

func (sb *StaticBuilder) Build(target resolver.Target, cc resolver.ClientConn,
	opts resolver.BuildOptions) (resolver.Resolver, error) {
	// use info in target to access naming service
	// parse the target.endpoint
	// resolver.Target{Scheme:"static", Authority:"", Endpoint:"localhost:5051,localhost:5052,localhost:5053"}
	endpoints := strings.Split(target.Endpoint, ",")
	log.Printf("static resolver: %s", target.Endpoint)
	for i, endpoint := range endpoints {
		if endpoint == ":" {
			endpoints[i] = "localhost:7070"
		} else if strings.HasPrefix(endpoint, ":") {
			endpoints[i] = "localhost" + endpoint
		}
	}
	log.Printf("static resolved: %v", endpoints)
	r := &StaticResolver{endpoints: endpoints, cc: cc}
	r.ResolveNow(resolver.ResolveNowOptions{})
	return r, nil
}

func (sb *StaticBuilder) Scheme() string { return "static" }

type StaticResolver struct {
	endpoints []string
	cc        resolver.ClientConn
	sync.Mutex
}

func (r *StaticResolver) ResolveNow(opts resolver.ResolveNowOptions) {
	r.Lock()
	r.doResolve()
	r.Unlock()
}

func (r *StaticResolver) Close() {}

func (r *StaticResolver) doResolve() {
	addrs := make([]resolver.Address, len(r.endpoints))
	for i, addr := range r.endpoints {
		addrs[i] = resolver.Address{
			Addr:       addr,
			ServerName: fmt.Sprintf("instance-%d", i+1),
		}
	}

	newState := resolver.State{Addresses: addrs}
	r.cc.UpdateState(newState)
}
