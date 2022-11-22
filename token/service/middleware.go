package service

import service "github.com/IRFAN374/upsvc/token"

type Middleware func(service.Service) service.Service
