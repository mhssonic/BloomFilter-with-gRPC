package request

import (
	"context"
	"hash/fnv"
)

const m uint64 = 350 * 1e7

var bitArray = make([]bool, m)

type ServerTemp interface {
}

type Server struct {
}

func (s *Server) mustEmbedUnimplementedBloomFilterServer() {
}

func (s *Server) AddString(ctx context.Context, message *Message) (*Message, error) {
	addString(message.Body)
	return &Message{Body: ""}, nil
}

func (s *Server) IsThere(ctx context.Context, msg *Message) (*Bool, error) {
	return &Bool{Body: bloomFilterIsThere(msg.Body)}, nil
}

func addString(str string) {
	i, j, k := hash(str)
	bitArray[i] = true
	bitArray[j] = true
	bitArray[k] = true
}

func bloomFilterIsThere(str string) bool {
	i, j, k := hash(str)
	return bitArray[i] && bitArray[j] && bitArray[k]
}

// return 3 hash code of a string
func hash(s string) (uint64, uint64, uint64) {
	h := fnv.New64a()
	h.Write([]byte(s))
	a := h.Sum64() % m
	h.Write([]byte(s))
	b := h.Sum64() % m
	h.Write([]byte(s))
	c := h.Sum64() % m
	return a, b, c
}
