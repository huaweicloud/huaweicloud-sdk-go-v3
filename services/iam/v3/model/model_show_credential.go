package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ShowCredential struct {
	// IAM用户ID。

	UserId string `json:"user_id"`
	// 查询的AK。

	Access string `json:"access"`
	// 访问密钥状态。

	Status string `json:"status"`
	// 访问密钥创建时间。

	CreateTime string `json:"create_time"`
	// 访问密钥的上次使用时间。

	LastUseTime string `json:"last_use_time"`
	// 访问密钥描述信息。

	Description string `json:"description"`
}

func (o ShowCredential) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCredential struct{}"
	}

	return strings.Join([]string{"ShowCredential", string(data)}, " ")
}
