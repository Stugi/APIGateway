package service

import (
	. "stugi/api-gateway/internal/model"
)

type Gateway interface {
	GetNews() ([]*NewsShortDetailed, error)
	FilterNews(filter string) ([]*NewsShortDetailed, error)
	GetNewsDetailed(id string) (*NewsFullDetailed, error)
	AddComment(comment *Comment) error
}

type gateway struct{}

func New() Gateway {
	return &gateway{}
}

func (g *gateway) GetNews() ([]*NewsShortDetailed, error) {
	return nil, nil
}

func (g *gateway) FilterNews(filter string) ([]*NewsShortDetailed, error) {
	return nil, nil
}

func (g *gateway) GetNewsDetailed(id string) (*NewsFullDetailed, error) {
	return nil, nil
}

func (g *gateway) AddComment(comment *Comment) error {
	return nil
}
