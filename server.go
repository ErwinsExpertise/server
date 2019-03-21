package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"net"

	"./packets"
)

func CheckError(err error) {
	if err != nil {
		log.Printf("Error occured: %+v", err)
	}
}

func main() {
	serv, err := net.Listen("tcp", "0.0.0.0:9000")
	CheckError(err)
	log.Println("Listening on :9000")
	//Begin server handling
	for {

		conn, err := serv.Accept()
		CheckError(err)
		fmt.Sprintf("[Address] %s", conn.RemoteAddr())
		defer conn.Close()

		go handleConnection(conn)
		continue
	}
}

func handleConnection(conn net.Conn) {
	for {
		data := make([]byte, 4)

		_, err := conn.Read(data)
		CheckError(err)

		pack := packets.NewPacket()
		pack.Append(data)

		fuck := packets.NewReader(&pack)
		log.Printf("\nReading: %+v ", fuck)
		process(conn, fuck)

	}

}

func readPacket(b []byte) (uint32, error) {
	log.Printf("\nReading: %+v ", b)
	if len(b) == 0 {
		return 0, errors.New("Invalid Response")
	} else {
		packet := binary.LittleEndian.Uint32(b[:])
		return packet, nil
	}
}

func sendPacket(b []byte, conn net.Conn) {
	_, err := conn.Write(b)
	CheckError(err)
}

func stringToByte(s string) []byte {
	var str = []string{s}
	var buf bytes.Buffer
	for _, s := range str {
		buf.WriteString(s)
	}
	return buf.Bytes()
}