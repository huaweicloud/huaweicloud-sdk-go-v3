package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Credentials struct {

	// IAM用户ID。
	UserId string `json:"user_id" xml:"user_id"`

	// 查询的AK。
	Access string `json:"access" xml:"access"`

	// 访问密钥状态。
	Status string `json:"status" xml:"status"`

	// 访问密钥创建时间。
	CreateTime string `json:"create_time" xml:"create_time"`

	// 访问密钥描述信息。
	Description string `json:"description" xml:"description"`
}

func (o Credentials) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Credentials struct{}"
	}

	return strings.Join([]string{"Credentials", string(data)}, " ")
}
