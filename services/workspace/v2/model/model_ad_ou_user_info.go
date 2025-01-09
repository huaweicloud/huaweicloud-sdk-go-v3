package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AdOuUserInfo struct {

	// 名称
	UserName *string `json:"user_name,omitempty"`

	// 过期时间 -1代表永不过期
	ExpiredTime *string `json:"expired_time,omitempty"`

	// 是否存在于用户列表
	HasExisted *bool `json:"has_existed,omitempty"`
}

func (o AdOuUserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdOuUserInfo struct{}"
	}

	return strings.Join([]string{"AdOuUserInfo", string(data)}, " ")
}
