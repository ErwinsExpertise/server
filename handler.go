package main

import (
	"net"

	"./packets"
)

func process(conn net.Conn, reader packets.Reader, key []byte) {
	switch reader.ReadByte() {
	case packets.Recv.DEF:
		resp := stringToByte("Default packet")
		sendPacket(resp, key, conn)
	case packets.Recv.LOGIN:
		resp := stringToByte("Login packet")
		sendPacket(resp, key, conn)
	case packets.Recv.MOV_LEFT:
		resp := stringToByte("Move left")
		sendPacket(resp, key, conn)
	case packets.Recv.MOV_RIGHT:
		resp := stringToByte("Move right")
		sendPacket(resp, key, conn)
	case packets.Recv.MOV_UP:
		resp := stringToByte("Move up")
		sendPacket(resp, key, conn)
	case packets.Recv.MOV_DOWN:
		resp := stringToByte("Move down")
		sendPacket(resp, key, conn)
	case packets.Recv.OP_EQUIP:
		resp := stringToByte("Open Equipment")
		sendPacket(resp, key, conn)
	default:
		resp := stringToByte("UNKNOWN")
		sendPacket(resp, key, conn)

	}
}
