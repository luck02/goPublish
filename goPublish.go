package goPublish

import ()

type PublishRequest struct {
}

type PublishMultipleRequest struct {
}

type PublishResponse struct {
}

type PublishMultipleResponse struct {
}

type Publisher interface {
	Publish(PublishRequest) PublishResponse
	PublishMultiple(PublishMultipleRequest) PublishMultipleResponse
}

type GoPublisher struct {
}

func NewGoPublisher() *GoPublisher {

	return &GoPublisher{}
}
