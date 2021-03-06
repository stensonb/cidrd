package plugin

// endpoints contains our registered endpoints
var endpoints []Endpoint

// RegisterEndpoint allows any plugin to register itself to the global
// endpoint registry.
func RegisterEndpoint(e Endpoint) {
	endpoints = append(endpoints, e)
}

// Endpoints will return the current list of endpoints
func Endpoints() []Endpoint {
	return endpoints
}
