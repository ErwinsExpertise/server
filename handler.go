package main

import (
	"net"

	"./packets"
)

func process(conn net.Conn, reader packets.Reader) {
	switch reader.ReadByte() {
	case packets.Recv.DEF:
		resp := stringToByte("Default packet")
		sendPacket(resp, conn)
	case packets.Recv.LOGIN:
		resp := stringToByte("Login packet")
		sendPacket(resp, conn)
	case packets.Recv.MOV_LEFT:
		resp := stringToByte("Move left bitch")
		sendPacket(resp, conn)
	case packets.Recv.MOV_RIGHT:
		resp := stringToByte("Move right bitch")
		sendPacket(resp, conn)
	case packets.Recv.MOV_UP:
		resp := stringToByte("Move up bitch")
		sendPacket(resp, conn)
	case packets.Recv.MOV_DOWN:
		resp := stringToByte("Move down bitch")
		sendPacket(resp, conn)
	case packets.Recv.OP_EQUIP:
		resp := stringToByte("Open Equipment bitch")
		sendPacket(resp, conn)
	default:
		resp := stringToByte("UNKNOWN")
		sendPacket(resp, conn)

	}
}
