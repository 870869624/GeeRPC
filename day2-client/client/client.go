package client

type Call struct {
	Seq           uint64
	ServiceMethod string
	Args          interface{}
	Replay        interface{}
	Error         error
	Done          chan *Call //Strobes when call is complete.
}

func (call *Call) done() {
	call.Done <- call
}

type Client struct {
	cc codec.Codec
}
