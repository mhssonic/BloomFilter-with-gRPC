package request

import (
	"BloomFilter"
	"context"
)

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
