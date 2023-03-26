package dto

import "gorm.io/gorm"

type Mock struct {
	Id          string          `json:"id"`
	Url         string          `json:"mainUrl"`
	Method      string          `json:"method"`
	Status      string          `json:"status"`
	ContentType string          `json:"contentType"`
	Headers		[]*Header		`json:"headers" gorm:"serializer:json"`
	Body        string          `json:"body"`
}
type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Mocks []*Mock

type LogRequest struct {
	gorm.Model
	MockId string `json:"mockId"`
	Body string `json:"body"`
}

type LogRequests []*LogRequest
