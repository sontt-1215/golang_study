package service

import (
)

type Office365Service interface {
	GetAuthCodeURL() string
}

type office365Service struct {
    
}

func NewOffice365Service() Office365Service {
    return &office365Service{}
}

func (office365 *office365Service) GetAuthCodeURL() string {
	return "hiiiii"
}
