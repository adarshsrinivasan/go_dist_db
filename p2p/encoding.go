package p2p

import (
	"encoding/gob"
	"io"
)

const BUFF_SIZE = 1028

type Decoder interface {
	Decode(io.Reader, *RPC) error
}

type GOBDecoder struct{}

func (g GOBDecoder) Decode(reader io.Reader, msg *RPC) error {
	return gob.NewDecoder(reader).Decode(msg)
}

type DefaultDecoder struct{}

func (g DefaultDecoder) Decode(reader io.Reader, msg *RPC) error {
	buf := make([]byte, BUFF_SIZE)
	n, err := reader.Read(buf)
	if err != nil {
		return err
	}
	msg.Payload = buf[:n]
	return nil

}
