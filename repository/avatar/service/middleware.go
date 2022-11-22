package service

import service "github.com/IRFAN374/upsvc/repository/avatar"

type Middleware func(service.Repository) service.Repository
