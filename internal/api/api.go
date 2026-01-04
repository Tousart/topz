package api

import "github.com/tousart/topz/internal/service"

type TopzApi struct {
	service *service.ProcService
}

func NewTopzApi(s *service.ProcService) *TopzApi {
	return &TopzApi{
		service: s,
	}
}
