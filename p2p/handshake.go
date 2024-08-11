package p2p

import "errors"

// Errors definition

// InvalidHandshake is returns when handshake between remote and local node fails
var InvalidHandshake = errors.New("invalid handshake")

type HandShakeFunc func(any) error

func NOPHandShakeFunc(any) error { return nil }
