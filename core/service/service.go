package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

type Context struct {
	w http.ResponseWriter
	r *http.Request
}

func (c *Context) Query(name string) string {
	return c.r.URL.Query().Get(name)
}

func (c *Context) Header(name string) string {
	return c.r.Header.Get(name)
}

func (c *Context) Read(v any) error {
	return json.NewDecoder(c.r.Body).Decode(&v)
}

func (c *Context) Write(v any) error {
	return json.NewEncoder(c.w).Encode(v)
}

func (c *Context) Status(status int) {
	c.w.WriteHeader(status)
}

func (c *Context) Error(status int, err any) error {

	c.w.Header().Add("Content-Type", "application/json")

	c.w.WriteHeader(status)

	var msg = ""

	if errMsg, isErr := err.(error); isErr {
		msg = errMsg.Error()
	} else if msgStr, isStr := err.(string); isStr {
		msg = msgStr
	} else {
		panic("invalid message type")
	}

	c.w.Write([]byte(
		fmt.Sprintf(`{"error":"%s"}`, msg),
	))

	return nil
}

type Handler func(Context)

type Service struct {
	Name    string
	Version string
	mux     chi.Router
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Service) Get(pattern string, h Handler)  { s.handle("GET", pattern, h) }
func (s *Service) Post(pattern string, h Handler) { s.handle("POST", pattern, h) }
func (s *Service) Put(pattern string, h Handler)  { s.handle("PUT", pattern, h) }

func (s *Service) handle(method, pattern string, h Handler) {
	s.mux.MethodFunc(method, pattern, func(w http.ResponseWriter, r *http.Request) {
		h(Context{w: w, r: r})
	})
}

func New(name, version string) (ant *Service) {

	srv := &Service{
		Name:    name,
		Version: version,
		mux:     chi.NewMux(),
	}

	if err := srv.Register(); err != nil {
		log.Printf("Failed to register service %s reason: %s\n", name, err.Error())
	}

	return srv
}

func (s *Service) Register() (err error) {

	buf := bytes.NewBufferString(fmt.Sprintf(`
		{
			"name":"%s",
			"addr":"%s"
		}
	`, s.Name, "http://localhost"+os.Getenv("LISTEN")))

	response, err := http.Post(os.Getenv("PROXY"), "application/json", buf)

	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusCreated {
		return errors.New(fmt.Sprintf("status code is %d", response.StatusCode))
	}

	return nil
}

func EndpointRequest(service, method, endpoint string, body io.Reader) (request *http.Request, err error) {

	proxy, exists := os.LookupEnv("PROXY")

	if !exists {
		return nil, errors.New("proxy not defined")
	}

	request, err = http.NewRequest(
		method,
		proxy+"/"+service+endpoint,
		body,
	)

	return
}
