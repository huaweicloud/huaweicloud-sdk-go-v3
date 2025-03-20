package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAccessKeyLastUsedV5Request Request Object
type ShowAccessKeyLastUsedV5Request struct {

	// IAM用户ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	UserId string `json:"user_id"`

	// 永久访问密钥ID，即AK。
	AccessKeyId string `json:"access_key_id"`
}

func (o ShowAccessKeyLastUsedV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccessKeyLastUsedV5Request struct{}"
	}

	return strings.Join([]string{"ShowAccessKeyLastUsedV5Request", string(data)}, " ")
}
