package packets

var Send send

type send struct {
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
	Send.DEF = 0x00
	Send.LOGIN = 0x01
	Send.MOV_LEFT = 0x02
	Send.MOV_RIGHT = 0x03
	Send.MOV_UP = 0x04
	Send.MOV_DOWN = 0x05
	Send.OP_EQUIP = 0x07
	Send.OP_STAT = 0x08
	Send.OP_FRIEND = 0x09
	Send.ATTACK = 0x0a
}
