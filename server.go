package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"log"
	"net"

	"./cipher"
	"./packets"
)

var key []byte

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
		log.Printf("Client connected: %s", conn.RemoteAddr())
		defer conn.Close()

		go handleConnection(conn)
		continue
	}
}

func handleConnection(conn net.Conn) {
	key = cipher.GenKey()
	_, err := conn.Write(key)
	CheckError(err)

	for {
		data := make([]byte, 512)

		id, err := conn.Read(data)
		if err != io.EOF {
			log.Printf("Connection closed: %s", conn.RemoteAddr())
			conn.Close()
		}

		packet := cipher.DecodeIV(key, data[:id])

		pack := packets.NewPacket()
		pack.Append(packet)

		load := packets.NewReader(&pack)
		log.Printf(" Client: %s Packet Recieved: %+v ", conn.RemoteAddr(), load)
		process(conn, load, key)

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

func sendPacket(b []byte, key []byte, conn net.Conn) {
	send := cipher.MakeIV(string(b), key)
	_, err := conn.Write(send)
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
