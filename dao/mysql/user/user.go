package user

import (
	"context"

	model "github.com/iooikaak/microService2/database/mysql/user"
)

//GetUserByID 根据一个ID获取单个User
func (d *Dao) GetUserByID(ctx context.Context, id int64) (result *model.User, err error) {
	result = &model.User{}
	err = d.Db.Context(ctx).Model(&model.User{}).Where("`user_id` = ?", id).First(result).Error
	return
}

//GetUserByIDs 根据多个ID批量获取Users
func (d *Dao) GetUserByIDs(ctx context.Context, ids []int64) (results []*model.User, err error) {
	results = make([]*model.User, 0)
	err = d.Db.Context(ctx).Model(&model.User{}).Where("`user_id` in (?)", ids).Order("create_time desc", true).Find(&results).Error
	return
}

//UpdateUser 更新User
func (d *Dao) UpdateUser(ctx context.Context, update *model.User) (err error) {
	err = d.Db.Context(ctx).Model(&model.User{}).Limit(1).Update(&update).Error
	return
}

//InsertUser 插入User
func (d *Dao) InsertUser(ctx context.Context, insert *model.User) (err error) {
	err = d.Db.Context(ctx).Omit("`user_id`").Create(&insert).Error
	return
}

func (d *Dao) GetUserInfo(ctx context.Context, userID int32) (result *model.UserInfo, err error) {
	result = &model.UserInfo{}
	err = d.Db.Context(ctx).Model(&model.UserInfo{}).Where("id = ?", userID).First(result).Error
	return
}
