package stress

import (
	"context"
	"fmt"
	_ "github.com/go-chassis/go-chassis/client/grpc"
	"github.com/go-chassis/go-chassis/client/rest"
	"github.com/go-chassis/go-chassis/core"
	"github.com/sirupsen/logrus"
	"github.com/tomlee0201/chassisdemo/protobuf"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

func SayRestHello(i int) bool {
	req, _ := rest.NewRequest("GET", "cse://RestServer/server/" + "user" + strconv.Itoa(i), nil)

	resp, err := core.NewRestInvoker().ContextDo(context.TODO(), req)
	if err != nil {
		logrus.Error(err)
		return false;
	}

	if resp.StatusCode != 200 {
		logrus.Info(resp.StatusCode)
		return false;
	} else {
		var responseBody []byte = make([]byte, resp.ContentLength)
		resp.Body.Read(responseBody)
		return true;
	}
}

func SayGRPCHello(i int) bool {
	reply := &protobuf.HelloReply{}
	//Invoke with microservice name, schema ID and operation ID
	if err := core.NewRPCInvoker().Invoke(context.Background(), "GRpcServer", "protobuf.Greeter", "SayHello",
		&protobuf.HelloRequest{Name: "user" + strconv.Itoa(i)}, reply, core.WithProtocol("grpc")); err != nil {
		logrus.Error("error" + err.Error())
		return false;
	} else {
		//logrus.Info(reply.Message)
		return true;
	}
}


func Run(args []string) {
	go func() {
		time.Sleep(1 * time.Second)

		startT := time.Now()
		var success int32 = 0
		var failure int32 = 0
		nthread := 100
		ncount := 1000
		for i := 0; i < nthread ; i++  {
			go func() {
				for j := 0; j < ncount; j++ {
					var result bool
					if len(args) == 0 || args[0] == "0" {
						result = SayGRPCHello(i)
					} else {
						result = SayRestHello(i)
					}
					if result {
						atomic.AddInt32(&success, 1)
					} else {
						atomic.AddInt32(&failure, 1)
					}
					time.Sleep(25 * time.Millisecond)
				}
			}()
		}
		for success + failure != int32(nthread * ncount) {
			 time.Sleep(100 * time.Millisecond)
		}

		fmt.Println("success %d, failure%d", success, failure)
		endT := time.Now()
		usedT := endT.Unix() - startT.Unix()
		fmt.Println("Use %ld sec", usedT)
	}()
}

func startProfile() {
	go func() {
		http.ListenAndServe("0.0.0.0:8089", nil)
	}()
}
