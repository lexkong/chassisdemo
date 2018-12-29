package restserver

import (
	"github.com/go-chassis/go-chassis"
	"github.com/go-chassis/go-chassis/server/restful"
	"net/http"
)


type RestFulHello struct {}



func (r *RestFulHello) Sayhello(b *restful.Context) {
	b.Write([]byte("Hello " + b.ReadPathParameter("userid") + " from rest server"))
}

func (s *RestFulHello) URLPatterns() []restful.Route {
	return []restful.Route{
		{Method:http.MethodGet, Path:"/server/{userid}", ResourceFuncName:"Sayhello"},
	}
}

func Run()  {
	chassis.RegisterSchema("rest", &RestFulHello{})
}

