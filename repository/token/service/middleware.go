package service

import service "github.com/IRFAN374/upsvc/repository/token"

type Middleware func(service.Repository) service.Repository
