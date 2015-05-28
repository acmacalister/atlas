package atlas

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
)

type RoundRobin struct {
	counter       int
	numOfBackends int
	downBackends  []string
}

func New(numOfBackends int) *RoundRobin {
	return &RoundRobin{counter: 0, numOfBackends: numOfBackends}
}

func (rr *RoundRobin) ServeLoadBalancer(w http.ResponseWriter, r *http.Request, backends []string) {
	be := backends[rr.counter]
	// need to check downBackends here.
	conn, _ := net.Dial("tcp", be) // need this to be settable, as opposed to hardset to tcp.

	buf := new(bytes.Buffer)
	r.Write(buf)
	io.Copy(conn, buf) // write the request to the backend connection.

	go func() {
		fmt.Println("up in here")
		writeBuf := new(bytes.Buffer)
		io.Copy(writeBuf, conn) // now copy the write back.
		fmt.Println("not going to make me go all out")
		fmt.Println(writeBuf)
	}()
	if rr.counter == rr.numOfBackends { // we are going to need some form of locking here.
		rr.counter++
	} else {
		rr.counter = 0
	}
	fmt.Println("I'm here")
}
