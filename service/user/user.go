package user

import (
	"context"
	"github.com/iooikaak/microService2/config"
	"github.com/iooikaak/microService2/dao/rocketmq/test"

	pbms2 "github.com/iooikaak/pb/microService2/http"

	model "github.com/iooikaak/microService2/database/mysql/user"
)

func (s *UserService) GetUserInfo(ctx context.Context, req *pbms2.GetUserInfoReq) (*model.UserInfo, error) {
	res, err := s.db.GetUserInfo(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	if err := test.New(config.Conf).SendToTest(ctx, "===== musheng2 =====", "", "test"); err != nil {
		return nil, err
	}
	return res, nil
}
