package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GrantUserInfo struct {

	// 资源ID
	UserId string `json:"user_id"`

	// 名称
	UserName string `json:"user_name"`

	// 创建时间
	CreateTime int64 `json:"create_time"`

	// 有效时间
	ValidityTime *int64 `json:"validity_time,omitempty"`
}

func (o GrantUserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GrantUserInfo struct{}"
	}

	return strings.Join([]string{"GrantUserInfo", string(data)}, " ")
}
