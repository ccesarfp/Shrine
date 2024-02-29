package application

import (
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"sync"
	"time"
)

type Application struct {
	Name    string
	Version string
	S       *server
	cb      *circuitBreaker
}

var (
	appOnce sync.Once
	i       *Application // instance
	kasp    = keepalive.ServerParameters{
		MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
		MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
		MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
	}
	errCh = make(chan error, 1)
)

func New() *Application {
	appOnce.Do(func() {
		i = &Application{
			S:  newServer(),
			cb: newCircuitBreaker(),
		}
	})
	return i
}

// Up - Start application
func (a *Application) Up() {
	// Writing gob
	err := write(i)
	if err != nil {
		log.Panicln(err)
	}

	log.Println("Starting listener")
	listener, err := net.Listen(i.S.Network, i.S.Address)
	if err != nil {
		_ = remove()
		log.Fatalln(err)
	}

	// Starting error recover function
	go i.errorRecover()

	// Starting server and starting channel
	log.Println("Application initialization took", time.Since(i.S.StartTime))
	errCh <- i.S.Server.Serve(listener)
	defer close(errCh)
}

// Down - shut down application
func (a *Application) Down() {
	log.Println("Shutting down application")
	// Removing gob
	err := remove()
	if err != nil {
		log.Println(err)
	}
	// Shutting down server
	i.S.Server.GracefulStop()
}

// DownBrutally - forcefully shutdown application
func (a *Application) DownBrutally() {
	log.Println("Brutally shutting down application")
	// Removing gob
	err := remove()
	if err != nil {
		log.Println(err)
	}
	// Shutting down server
	i.S.Server.Stop()
}

// errorRecover - if have errors, increment variable. If have more or exactly 3 errors, stop server
func (a *Application) errorRecover() {
	for {
		select {
		// an error has occurred, so increment variable and set error time
		case <-errCh:
			i.cb.incrementError()
			if i.cb.quantity >= 3 {
				a.DownBrutally()
				log.Fatalf("Shutting down server. Error happened 3 times or more.")
			}
		// every 10 seconds check if any error occurred and 5 minutes have passed, so clear the errors and error time
		case <-time.After(10 * time.Second):
			i.cb.verifyError()
		}
	}
}
