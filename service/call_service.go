package service

import (
	"gopkg.in/resty.v1"
	"time"
)

type callService struct {
}

// NewCallService implements the CallService Interface
func NewCallService() CallService {
	return &callService{}
}

type CallService interface {
	Request() *resty.Request
}

func (c *callService) Request() *resty.Request {
	return resty.New().
		RemoveProxy().
		SetDebug(false).
		SetRedirectPolicy(resty.FlexibleRedirectPolicy(15)).
		SetTimeout(180 * time.Second).
		R()
}
