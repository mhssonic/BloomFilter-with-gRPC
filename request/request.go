package request

import (
	"BloomFilter"
	"context"
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
	BloomFilter.AddString(message.Body)
	return &Message{Body: ""}, nil
}

func (s *Server) IsThere(ctx context.Context, msg *Message) (*Bool, error) {
	return &Bool{Body: BloomFilter.IsThere(msg.Body)}, nil
}
