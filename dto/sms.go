package dto

import "github.com/Hot-One/kizen-go-service/pkg/pg"

type SmsPage = pg.PageData[Sms] // @name SmsPage

type Sms struct {
	Id        int64  `json:"id"`
	Type      string `json:"type"`
	Value     string `json:"value"`
	SentCount int    `json:"sent_count"`
	SentAt    string `json:"sent_at"`
	Code      string `json:"code"`
	CreatedAt string `json:"created_at" swaggerignore:"true"`
	UpdatedAt string `json:"updated_at" swaggerignore:"true"`
} // @name Sms

type CreateSms struct {
	Type  string `json:"type" binding:"required"`
	Value string `json:"value" binding:"required"`
	Code  string `json:"code" binding:"required" swaggerignore:"true"`
} // @name SendSms

type UpdateSms struct {
	SentCount *int `json:"sent_count"`
	SentAt    *int `json:"sent_at"`
} // @name UpdateSms

type VerifySms struct {
	Id   int64  `json:"id" binding:"required" form:"id" query:"id"`
	Code string `json:"code" binding:"required" form:"code" query:"code"`
} // @name VerifySms
