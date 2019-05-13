package tmsg

type MType uint64

type Header struct {
	Type MType
}

func (h *Header) GetType() MType {
	return h.Type
}

type IsTMsg interface {
	GetType() MType
}
