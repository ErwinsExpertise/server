package packets

var Recv recv

type recv struct {
	DEF       byte
	LOGIN     byte
	MOV_LEFT  byte
	MOV_RIGHT byte
	MOV_UP    byte
	MOV_DOWN  byte
	OP_EQUIP  byte
	OP_STAT   byte
	OP_FRIEND byte
	ATTACK    byte
}

func init() {
	Recv.DEF = 0x00
	Recv.LOGIN = 0x01
	Recv.MOV_LEFT = 0x02
	Recv.MOV_RIGHT = 0x03
	Recv.MOV_UP = 0x04
	Recv.MOV_DOWN = 0x05
	Recv.OP_EQUIP = 0x07
	Recv.OP_STAT = 0x08
	Recv.OP_FRIEND = 0x09
	Recv.ATTACK = 0x0a
}
