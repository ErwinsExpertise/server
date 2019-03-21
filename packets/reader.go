package packets

// Reader -
type Reader struct {
	pos    int
	packet *Packet
}

func NewReader(p *Packet) Reader {
	return Reader{pos: 0, packet: p}
}

func (r *Reader) ReadByte() byte {
	return r.packet.readByte(&r.pos)
}
