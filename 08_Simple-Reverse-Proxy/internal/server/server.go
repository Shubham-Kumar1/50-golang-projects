package server

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/Shubham-Kumar1/08_Simple-Reverse-Proxy/internal/configs"
)

// Run starts server and listens on defined port
func Run() error {
	// load configurations from config file
	config, err := configs.NewConfiguration()
	if err != nil {
		return fmt.Errorf("could not load configuration: %v", err)
	}
	// Creates a new router
	mux := http.NewServeMux()
	// Registering the healthcheck endpoint
	mux.HandleFunc("/ping", ping)
	// Iterating through the configuration resource and registering them
	// into the router.
	for _, resource := range config.Resources {
		url, err := url.Parse(resource.Destination_URL)
		if err != nil {
			return fmt.Errorf("invalid destination URL for resource %s: %v", resource.Name, err)
		}
		proxy := NewProxy(url)
		mux.HandleFunc(resource.Endpoint, ProxyRequestHandler(proxy, url, resource.Endpoint))
	}
	// Running proxy server
	if err := http.ListenAndServe(config.Server.Host+":"+config.Server.Listen_port, mux); err != nil {
		return fmt.Errorf("could not start the server: %v", err)
	}
	return nil
}
