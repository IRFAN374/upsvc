package service

import service "github.com/IRFAN374/upsvc/user"

type Middleware func(service.Service) service.Service
