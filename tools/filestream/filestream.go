package filestream

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

type FileServer struct {
	port int
}

func NewFileServer(port int) *FileServer {

	return &FileServer{
		port: port,
	}
}

func (fs *FileServer) start() {
	// initiates tcp server listener
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", fs.port))

	if err != nil {
		log.Fatal("Server could not start.")
	}

	// continuosly listen to connections and spin up goroutines

	for {
		conn, err := ln.Accept() // accept and establish connection

		if err != nil {
			log.Fatal("Connection could not be established.")
		}

		go fs.serverReadLoop(conn)
	}

}

func (fs *FileServer) serverReadLoop(conn net.Conn) {
	buf := make([]byte, 2048) // 2 KB

	for {
		// read data from the connection - execution is blocked until data is received
		n, err := conn.Read(buf) // slices are reference types, passing it automatically pass reference

		if err != nil {
			log.Fatal("Error when reading message to server.")
		}

		// use the stored portion (buffer not filled up)
		file := buf[:n]
		fmt.Println("file recieved:", file)

		// convert file back to values
		fileContents := string(file)
		fmt.Printf("file contents (length: %d) decoded to utf-8:\n\n %s", n, fileContents)
	}
}

// simulate sending file
func sendFile(address int) error {
	file, err := os.Open("tools/filestream/data/test-file.txt")

	if err != nil {
		return err
	}

	fileContents, err := io.ReadAll(file)

	if err != nil {
		return err
	}

	// connect to the tcp server
	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", address))

	if err != nil {
		return err
	}

	// send file
	_, err = conn.Write(fileContents)

	if err != nil {
		return err
	}

	return nil
}

func RunFileStream() {
	address := 3232
	go func() {

		time.Sleep(time.Millisecond * 1500)
		// send file
		err := sendFile(address)

		if err != nil {
			fmt.Printf("Error when sending file %s", err)
		}
	}()

	fs := NewFileServer(address)
	// start server
	fs.start()
}
