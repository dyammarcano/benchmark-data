package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type PingPongService struct {
	buffer        []byte
	maxBufferSize int
	count         int
	mutex         sync.Mutex
}

func NewPingPongService(maxBufferSize int) *PingPongService {
	service := &PingPongService{
		maxBufferSize: maxBufferSize,
	}
	go service.startCounting()
	return service
}

func (s *PingPongService) startCounting() {
	go func() {
		for s.count <= 20000 {
			s.mutex.Lock()
			if s.buffer == nil || len(s.buffer) >= s.maxBufferSize {
				if s.buffer != nil {
					s.buffer = nil
				}
				s.buffer = make([]byte, 1*1024*1024) // 1MB
			}
			s.buffer = append(s.buffer, make([]byte, 1*1024*1024)...)
			s.count++
			s.mutex.Unlock()
			time.Sleep(10 * time.Millisecond)
		}
	}()
}

func (s *PingPongService) GetCount() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.count
}

type HomeController struct {
	pingPongService *PingPongService
}

func NewHomeController(pingPongService *PingPongService) *HomeController {
	return &HomeController{pingPongService: pingPongService}
}

func (controller *HomeController) home(w http.ResponseWriter, r *http.Request) {
	count := controller.pingPongService.GetCount()
	fmt.Fprintf(w, "Total ping-pong count: %d\n", count)
}

func (controller *HomeController) hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Docker, from Go App!")
}

func main() {
	pingPongService := NewPingPongService(1000 * 1024 * 1024)
	homeController := NewHomeController(pingPongService)

	http.HandleFunc("/", homeController.home)
	http.HandleFunc("/hello", homeController.hello)

	fmt.Println("Server is running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
