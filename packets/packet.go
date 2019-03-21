package packets

type Packet []byte

func NewPacket() Packet {
	return make(Packet, 0)
}

func CreateWithOpCode(op byte) Packet {
	p := Packet{}
	p.WriteInt32(0)
	p.WriteByte(op)

	return p
}

func (p *Packet) Append(data []byte) {
	*p = append(*p, data...)
}
func (p *Packet) WriteByte(data byte) {
	*p = append(*p, data)
}

func (p *Packet) WriteUint32(data uint32) {
	*p = append(*p, byte(data), byte(data>>8), byte(data>>16), byte(data>>24))
}

func (p *Packet) readUint32(pos *int) uint32 {
	return uint32(p.readByte(pos)) |
		uint32(p.readByte(pos))<<8 |
		uint32(p.readByte(pos))<<16 |
		uint32(p.readByte(pos))<<24
}

func (p *Packet) WriteInt32(data int32) { p.WriteUint32(uint32(data)) }

func (p *Packet) readByte(pos *int) byte {
	r := byte((*p)[*pos])
	*pos++
	return r
}

func (p *Packet) readBytes(pos *int, length int) []byte {
	r := []byte((*p)[*pos : *pos+length])
	*pos += length
	return r
}
