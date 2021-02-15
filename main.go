package main 

import ( 
	"fmt"
	"github.com/kardianos/service"
	"time"
)

const serviceName = "Medium Service"
const serviceDescription = "Simple service, just for fun"

type program struct {}

func (p program) Start(s service.Service) error {
	fmt.Println(s.String() + " started")
	go p.run()
	// panic("implement me")
	return nil
}

func (p program) Stop(s service.Service) error {
	// panic("implement me")
	fmt.Println(s.String() + " stopped")
	return nil
}

func (p program) run() {
	for {
		fmt.Println("Service is running")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	serviceConfig := &service.Config{
		Name: serviceName,
		DisplayName: serviceName,
		Description: serviceDescription,
	}
	prg := &program{}
	s, err := service.New(prg, serviceConfig)
	if err != nil {
		fmt.Println("Cannot create the service: " + err.Error())
	}
	err = s.Run()
	if err != nil {
		fmt.Println("Cannot start the service: " + err.Error())
	}
}

// var (
// 	serviceIsRunning bool
// 	programming bool
// 	writingSync sync.Mutex
// )