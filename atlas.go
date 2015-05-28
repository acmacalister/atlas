package atlas

import "net/http"

type server struct {
	backends []string
	balance  Balancer
}

type Balancer interface {
	ServeLoadBalancer(w http.ResponseWriter, r *http.Request, backends []string)
}

type Frontend struct {
	// Host should be a resolvable host that can be connected to.
	// This is usuallythe same string address you would pass to http.ListenAndServe.
	Host string

	// SSLEnabled is a property to set if you should are going to do SSL Termination with this frontend.
	SSLEnabled bool

	// Cert is the path to your SSL cert. Only needed if SSLEnabled is set to true.
	Cert string

	// Key is the path to your SSL key. Only needed if SSLEnabled is set to true.
	Key string
}

func Run(frontends []Frontend, backends []string, balance Balancer) error {
	s := server{balance: balance, backends: backends}
	http.HandleFunc("/", s.handler)

	length := len(frontends)
	for i, fe := range frontends {
		if i+1 == length {
			if err := startFrontend(&fe); err != nil {
				return err
			}
		} else {
			go startFrontend(&fe) // need a good way to handle these errors...
		}
	}

	return nil // should never get here.
}

func (s *server) handler(w http.ResponseWriter, r *http.Request) {
	s.balance.ServeLoadBalancer(w, r, s.backends)
}

func startFrontend(fe *Frontend) error {
	if fe.SSLEnabled {
		return http.ListenAndServeTLS(fe.Host, fe.Cert, fe.Key, nil)
	}
	return http.ListenAndServe(fe.Host, nil)
}
