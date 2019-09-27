package server

import (
	context "context"

	"github.com/{{.Org}}/{{.Proj}}/services/greetings"

	"github.com/{{.Org}}/{{.Proj}}/internal/config"
	"go.uber.org/zap"
)

// Server represents the gRPC server
type Server struct {
	Log *zap.SugaredLogger
	Cfg *config.Config
}

// Greet gets a product from manifold
func (s *Server) Greet(ctx context.Context, in *greetings.GreetReq) (*greetings.GreetResp, error) {
	res := &greetings.GreetResp{
		Response: &greetings.Greeting{
			Msg: "Welcome to the future!",
		},
	}

	return res, nil
}
