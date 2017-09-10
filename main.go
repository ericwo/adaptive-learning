package main

import (
	"fmt"
	"net/http"

	"github.com/ericwo/adaptive-learning/config"
	"github.com/theplant/appkit/server"
)

func main() {
	mux := http.NewServeMux()

	httpLogger := config.Logger.With("origin", "http")
	httpLogger.Debug().Log("msg", "Initializing http.Handler instance...")

	middleware := server.Compose(
		server.DefaultMiddleware(httpLogger),
	)

	port := fmt.Sprintf(":%d", config.Config.Port)
	fmt.Println("listen on", port)

	http.ListenAndServe(port, middleware(mux))
}
