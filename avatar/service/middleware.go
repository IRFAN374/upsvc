package service

import service "github.com/IRFAN374/upsvc/avatar"

type Middleware func(service.Service) service.Service
