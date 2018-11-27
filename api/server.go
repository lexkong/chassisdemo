package api

import (
	"github.com/go-chassis/go-chassis/server/restful"
	"github.com/go-chassis/go-chassis/core/lager"
	"github.com/go-chassis/go-chassis/client/rest"
	"github.com/go-chassis/go-chassis/core"
	"github.com/sirupsen/logrus"
	_ "github.com/go-chassis/go-chassis/client/grpc"
	"net/http"
	"github.com/go-chassis/go-chassis"
	"github.com/tomlee0201/chassisdemo/protobuf"
	"context"
)

type RestFulApi struct {}

func (r *RestFulApi) SayRestHello(b *restful.Context) {
	lager.Logger.Infof("Request:%s", b.ReadPathParameter("userid"))

	req, _ := rest.NewRequest("GET", "cse://RestServer/server/" + b.ReadPathParameter("userid"), nil)

	resp, err := core.NewRestInvoker().ContextDo(context.TODO(), req)
	if err != nil {
		b.WriteHeader(http.StatusServiceUnavailable)
		b.Write([]byte("Server internal error"))
		logrus.Error(err)
		return
	}

	if resp.StatusCode != 200 {
		b.WriteHeader(resp.StatusCode)
		b.Write([]byte("Server internal error"))
	} else {
		var responseBody []byte = make([]byte, resp.ContentLength)
		resp.Body.Read(responseBody)
		logrus.Info(string(responseBody))
		b.Write(responseBody)
	}
}

func (r *RestFulApi) SayGRPCHello(b *restful.Context) {
	lager.Logger.Infof("Request:%s", b.ReadPathParameter("userid"))

	//declare reply struct
	reply := &protobuf.HelloReply{}
	//Invoke with microservice name, schema ID and operation ID
	if err := core.NewRPCInvoker().Invoke(context.Background(), "GRpcServer", "helloworld.Greeter", "SayHello",
		&protobuf.HelloRequest{Name: b.ReadPathParameter("userid")}, reply, core.WithProtocol("grpc")); err != nil {
		logrus.Error("error" + err.Error())
		b.WriteHeader(http.StatusInternalServerError)
		b.Write([]byte("Server internal error"))
	} else {
		logrus.Info(reply.Message)
		b.Write([]byte(reply.Message))
	}
}

func (s *RestFulApi) URLPatterns() []restful.Route {
	return []restful.Route{
		{http.MethodGet, "/sayresthello/{userid}", "SayRestHello"},
		{http.MethodGet, "/saygrpchello/{userid}", "SayGRPCHello"},
	}
}

func Run() {
	//start all server you register in server/schemas.
	chassis.RegisterSchema("rest", &RestFulApi{})

}

