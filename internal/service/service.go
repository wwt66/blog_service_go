package service

import (
	"context"
	"myblogservice/global"
	"myblogservice/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine) // todo 这里确认一下DB是否创建了新的
	return svc
}
