package naming

import (
	clientProxy "github.com/arma29/mid-rpc/my-middleware/distribution/clientProxy"
)

type NamingService struct {
	Repository map[string]clientProxy.ClientProxy
}

func (naming *NamingService) Register(name string, proxy clientProxy.ClientProxy) (bool) {
	r := false

	// check if repository is already created
	if len(naming.Repository) == 0 {
		naming.Repository = make(map[string]clientProxy.ClientProxy)
	}
	// check if the service is already registered
	_, ok := naming.Repository[name]
	if ok {
		r = false // service already registered
	} else { // service not registered
		naming.Repository[name] = clientProxy.ClientProxy{Host: proxy.Host, Port: proxy.Port}
		r = true
	}

	return r
}

func (naming NamingService) Lookup(name string) clientProxy.ClientProxy {
	return naming.Repository[name]
}

