package concurrency

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

/*
Task: If there are 1000 requests in which there are only 100 unique user ids then there should be only a maximum 100
requests **into the database** but all 1000 requests should get a response with a user data.
*/

func RunRequestThrottle() {
	fmt.Println("Request Throttle Challenge")
	fmt.Println("--------------------------")

	mux := http.NewServeMux()

	// server
	serverOpts := NewServer(":9999")

	// routes
	mux.HandleFunc("GET /api/users/{id}", serverOpts.userHandler)

	go func() {
		time.Sleep(time.Second * 2)

		serverOpts.testRequests()
	}()

	server := &http.Server{
		Addr:    serverOpts.addr,
		Handler: mux,
	}

	fmt.Println("Server listening on ", server.Addr)
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Server could not start due to error:", err)
	}
}

type DBSchema map[uint]User

type User struct {
	ID    uint
	Name  string
	Email string
}

type ServerOpts struct {
	DB    DBSchema
	Cache DBSchema
	addr  string
}

const dbMaxSize = 100

func NewServer(addr string) *ServerOpts {
	dbDefaultData := make(DBSchema, dbMaxSize) // NOTE: sets capacity, NOT ACTUAL length

	for i := 0; i < dbMaxSize; i++ {
		index := uint(i)

		dbDefaultData[index] = User{
			ID:    index + 1,
			Name:  fmt.Sprintf("User%d", index+1),
			Email: fmt.Sprintf("user%d@test.com", index+1),
		}
	}

	return &ServerOpts{
		DB:    dbDefaultData,
		Cache: make(DBSchema, 100),
		addr:  addr,
	}
}

type Req struct {
	ID uint `json:"id"`
}

// caching layer
func (s *ServerOpts) cacheRequest(id uint) bool {
	_, ok := s.Cache[id]

	if !ok {
		// update cache if that user doesn't exist there already
		s.Cache[id] = s.DB[id]
	}

	return ok
}

func (s *ServerOpts) userHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	idInt, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		fmt.Println("Error when converting id to int:", err)
	}

	index := uint(idInt) - 1

	// check if data exists in cache first
	var userData User
	exists := s.cacheRequest(index)
	if !exists {
		fmt.Printf("Getting from DB user with id %d\n", index)
		userData = s.DB[index]
	} else {
		fmt.Printf("Getting from Cache user with id %d\n", index)
		userData = s.Cache[index]
	}

	if index < 0 {
		http.Error(w, "Index cannot be less than 1.", http.StatusBadRequest)
		return
	}

	out, err := json.Marshal(userData)

	w.Write(out)
}

func (s *ServerOpts) testRequests() {

	for id := 0; id < 3; id++ {

		endpoint := fmt.Sprintf("http://localhost%s/api/users/47", s.addr)
		fmt.Println("Endpoint:", endpoint)

		resp, err := http.Get(endpoint)

		if err != nil {
			fmt.Printf("Error when making the GET request: %v", err)
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error when reading the response body: %v", err)
		}

		// Print the response
		fmt.Println("Response:", string(body))
	}
}
