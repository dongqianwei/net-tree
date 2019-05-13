package tmsg

type Processor struct {
	handlers      map[MType]Handler
	defaultHandle Handler
	input         <-chan IsTMsg
}

type Handler interface {
	Handle(msg IsTMsg)
}

func (p *Processor) RegisterHandler(mType MType, handler Handler) {
	if _, ok := p.handlers[mType]; ok {
		panic("already registered=" + string(mType))
	}

	p.handlers[mType] = handler
}

func (p *Processor) process() {
	for {
		msg := <-p.input
		mType := msg.GetType()
		if handler, ok := p.handlers[mType]; ok {
			handler.Handle(msg)
		} else {
			if p.defaultHandle != nil {
				p.defaultHandle.Handle(msg)
			}
		}
	}
}
