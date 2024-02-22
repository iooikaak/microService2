package model

import (
	"github.com/iooikaak/frame/util"
)

// User 用户表
type User struct {
	ID            int64         `gorm:"primaryKey;column:id" json:"-"`               // 自增ID
	UserID        int64         `gorm:"column:user_id" json:"userId"`                // 用户id
	Username      string        `gorm:"column:username" json:"username"`             // 会员用户名
	Password      string        `gorm:"column:password" json:"password"`             // 会员密码
	Mobile        string        `gorm:"column:mobile" json:"mobile"`                 // 手机号码
	Face          string        `gorm:"column:face" json:"face"`                     // 会员头像
	Nickname      string        `gorm:"column:nickname" json:"nickname"`             // 会员昵称
	Sex           int8          `gorm:"column:sex" json:"sex"`                       // 会员性别 0男 1女
	Birthday      util.JsonTime `gorm:"column:birthday" json:"birthday"`             // 用户生日
	ClientType    int8          `gorm:"column:client_type" json:"clientType"`        // 客户端 0 小程序
	Province      int64         `gorm:"column:province" json:"province"`             // 省
	City          int64         `gorm:"column:city" json:"city"`                     // 市
	Area          int64         `gorm:"column:area" json:"area"`                     // 区
	Address       string        `gorm:"column:address" json:"address"`               // 详细地址
	CreateUID     int64         `gorm:"column:create_uid" json:"createUid"`          // 创建者
	UpdateUID     int64         `gorm:"column:update_uid" json:"updateUid"`          // 更新者
	Disabled      int64         `gorm:"column:disabled" json:"disabled"`             // 会员状态 true/false 禁用/未禁用
	IsDeleted     int64         `gorm:"column:is_deleted" json:"isDeleted"`          // 删除标志 true/false 删除/未删除
	LastLoginTime util.JsonTime `gorm:"column:last_login_time" json:"lastLoginTime"` // 最后一次登录时间
	CreateTime    util.JsonTime `gorm:"column:create_time" json:"createTime"`        // 创建时间
	UpdateTime    util.JsonTime `gorm:"column:update_time" json:"updateTime"`        // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *User) TableName() string {
	return "tb_user"
}

// UserColumns get sql column name.获取数据库列名
var UserColumns = struct {
	ID            string
	UserID        string
	Username      string
	Password      string
	Mobile        string
	Face          string
	Nickname      string
	Sex           string
	Birthday      string
	ClientType    string
	Province      string
	City          string
	Area          string
	Address       string
	CreateUID     string
	UpdateUID     string
	Disabled      string
	IsDeleted     string
	LastLoginTime string
	CreateTime    string
	UpdateTime    string
}{
	ID:            "id",
	UserID:        "user_id",
	Username:      "username",
	Password:      "password",
	Mobile:        "mobile",
	Face:          "face",
	Nickname:      "nickname",
	Sex:           "sex",
	Birthday:      "birthday",
	ClientType:    "client_type",
	Province:      "province",
	City:          "city",
	Area:          "area",
	Address:       "address",
	CreateUID:     "create_uid",
	UpdateUID:     "update_uid",
	Disabled:      "disabled",
	IsDeleted:     "is_deleted",
	LastLoginTime: "last_login_time",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// UserCollect 买家用户收藏表
type UserCollect struct {
	ID          int64         `gorm:"primaryKey;column:id" json:"-"`          // 自增ID
	UserID      int64         `gorm:"column:user_id" json:"userId"`           // 用户表id
	StoreuserID int64         `gorm:"column:storeuser_id" json:"storeuserId"` // 店铺用户id(主键)
	GoodsID     int64         `gorm:"column:goods_id" json:"goodsId"`         // 商品id
	MerchantID  int64         `gorm:"column:merchant_id" json:"merchantId"`   // 商家id
	IsDeleted   int64         `gorm:"column:is_deleted" json:"isDeleted"`     // 是否取消收藏 true/false
	CreateTime  util.JsonTime `gorm:"column:create_time" json:"createTime"`   // 创建时间
	UpdateTime  util.JsonTime `gorm:"column:update_time" json:"updateTime"`   // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *UserCollect) TableName() string {
	return "tb_user_collect"
}

// UserCollectColumns get sql column name.获取数据库列名
var UserCollectColumns = struct {
	ID          string
	UserID      string
	StoreuserID string
	GoodsID     string
	MerchantID  string
	IsDeleted   string
	CreateTime  string
	UpdateTime  string
}{
	ID:          "id",
	UserID:      "user_id",
	StoreuserID: "storeuser_id",
	GoodsID:     "goods_id",
	MerchantID:  "merchant_id",
	IsDeleted:   "is_deleted",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
}

// UserFavor 买家用户点赞表
type UserFavor struct {
	ID          int64         `gorm:"primaryKey;column:id" json:"-"`          // 自增ID
	UserID      int64         `gorm:"column:user_id" json:"userId"`           // 用户表id
	StoreuserID int64         `gorm:"column:storeuser_id" json:"storeuserId"` // 店铺用户id
	GoodsID     int64         `gorm:"column:goods_id" json:"goodsId"`         // 商品id
	MerchantID  int64         `gorm:"column:merchant_id" json:"merchantId"`   // 商家id
	IsDeleted   int64         `gorm:"column:is_deleted" json:"isDeleted"`     // 是否取消点赞 true/false
	CreateTime  util.JsonTime `gorm:"column:create_time" json:"createTime"`   // 创建时间
	UpdateTime  util.JsonTime `gorm:"column:update_time" json:"updateTime"`   // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *UserFavor) TableName() string {
	return "tb_user_favor"
}

// UserFavorColumns get sql column name.获取数据库列名
var UserFavorColumns = struct {
	ID          string
	UserID      string
	StoreuserID string
	GoodsID     string
	MerchantID  string
	IsDeleted   string
	CreateTime  string
	UpdateTime  string
}{
	ID:          "id",
	UserID:      "user_id",
	StoreuserID: "storeuser_id",
	GoodsID:     "goods_id",
	MerchantID:  "merchant_id",
	IsDeleted:   "is_deleted",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
}

// UserOperateLog 用户模块-用户操作记录表
type UserOperateLog struct {
	ID            int64         `gorm:"primaryKey;column:id" json:"-"`               // 自增ID
	OperatorLogID int64         `gorm:"column:operator_log_id" json:"operatorLogId"` // 操作记录表id
	UserID        int64         `gorm:"column:user_id" json:"userId"`                // 用户表id
	OperateType   string        `gorm:"column:operate_type" json:"operateType"`      // 操作类型 LOGIN-登录，ORDER-下单，COLLECT-收藏，FAVOR-点赞等
	OperateDesc   string        `gorm:"column:operate_desc" json:"operateDesc"`      // 操作描述
	IsDeleted     int64         `gorm:"column:is_deleted" json:"isDeleted"`          // 删除标志 true/false 删除/未删除
	CreateTime    util.JsonTime `gorm:"column:create_time" json:"createTime"`        // 创建时间
	UpdateTime    util.JsonTime `gorm:"column:update_time" json:"updateTime"`        // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *UserOperateLog) TableName() string {
	return "tb_user_operate_log"
}

// UserOperateLogColumns get sql column name.获取数据库列名
var UserOperateLogColumns = struct {
	ID            string
	OperatorLogID string
	UserID        string
	OperateType   string
	OperateDesc   string
	IsDeleted     string
	CreateTime    string
	UpdateTime    string
}{
	ID:            "id",
	OperatorLogID: "operator_log_id",
	UserID:        "user_id",
	OperateType:   "operate_type",
	OperateDesc:   "operate_desc",
	IsDeleted:     "is_deleted",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}

// UserUpgradeStoreuserApproveLog 个人买家升级成店铺买家审批记录表
type UserUpgradeStoreuserApproveLog struct {
	ID            int64         `gorm:"primaryKey;column:id" json:"-"`              // 自增ID
	ApproveID     int64         `gorm:"column:approve_id" json:"approveId"`         // 审批表id
	UserID        int64         `gorm:"column:user_id" json:"userId"`               // 用户表id
	ApptitudePic1 string        `gorm:"column:apptitude_pic1" json:"apptitudePic1"` // 资质图片
	ApptitudePic2 string        `gorm:"column:apptitude_pic2" json:"apptitudePic2"` // 资质图片
	ApptitudePic3 string        `gorm:"column:apptitude_pic3" json:"apptitudePic3"` // 资质图片
	OperatorID    int64         `gorm:"column:operator_id" json:"operatorId"`       // 审批人id
	OperatorName  string        `gorm:"column:operator_name" json:"operatorName"`   // 审批人姓名
	IsApprove     int64         `gorm:"column:is_approve" json:"isApprove"`         // 是否审批通过 true/false
	Reason        string        `gorm:"column:reason" json:"reason"`                // 通过/不通过原因
	IsDeleted     int64         `gorm:"column:is_deleted" json:"isDeleted"`         // 软删除 true/false
	CreateTime    util.JsonTime `gorm:"column:create_time" json:"createTime"`       // 创建时间
	UpdateTime    util.JsonTime `gorm:"column:update_time" json:"updateTime"`       // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *UserUpgradeStoreuserApproveLog) TableName() string {
	return "tb_user_upgrade_storeuser_approve_log"
}

// UserUpgradeStoreuserApproveLogColumns get sql column name.获取数据库列名
var UserUpgradeStoreuserApproveLogColumns = struct {
	ID            string
	ApproveID     string
	UserID        string
	ApptitudePic1 string
	ApptitudePic2 string
	ApptitudePic3 string
	OperatorID    string
	OperatorName  string
	IsApprove     string
	Reason        string
	IsDeleted     string
	CreateTime    string
	UpdateTime    string
}{
	ID:            "id",
	ApproveID:     "approve_id",
	UserID:        "user_id",
	ApptitudePic1: "apptitude_pic1",
	ApptitudePic2: "apptitude_pic2",
	ApptitudePic3: "apptitude_pic3",
	OperatorID:    "operator_id",
	OperatorName:  "operator_name",
	IsApprove:     "is_approve",
	Reason:        "reason",
	IsDeleted:     "is_deleted",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
}
