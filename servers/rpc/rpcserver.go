package main

import (
	"fmt"
	"github.com/curso/structs_interfaces/structs"
	"log"
	"math"
	"net"
	"net/http"
	"net/rpc"
)

type AnotherCircle structs.Circle

func (c *AnotherCircle) SetFields(fields ...float64) {
	c.X = fields[0]
	c.Y = fields[1]
	c.R = fields[2]
}

func (c AnotherCircle) RemoteArea(fields *AnotherCircle, ret *float64) error {
	*ret = math.Pi * fields.R * fields.R
	return nil
}

func server() {
	circle := new(AnotherCircle)
	rpc.Register(circle)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}
func client() error {
	circle := new(AnotherCircle)
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
		return err
	}
	// Synchronous call
	circle.SetFields(0, 0, 10)
	var reply float64
	err = client.Call("AnotherCircle.RemoteArea", &circle, &reply)
	if err != nil {
		log.Fatal("Circle error:", err)
		return err
	}
	fmt.Printf("Circle area: %f\n", reply)
	// Asynchronous call
	areaCall := client.Go("AnotherCircle.RemoteArea", &circle, &reply, nil)
	replyCall := <-areaCall.Done
	fmt.Printf("Circle area: %f\n", *replyCall.Reply.(*float64))
	return nil
}

func main() {
	server()
	client()
}
