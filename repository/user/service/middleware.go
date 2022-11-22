package service

import service "github.com/IRFAN374/upsvc/repository/user"

type Middleware func(service.Repository) service.Repository
