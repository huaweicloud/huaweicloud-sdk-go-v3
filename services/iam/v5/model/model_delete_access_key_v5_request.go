package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAccessKeyV5Request Request Object
type DeleteAccessKeyV5Request struct {

	// IAM用户ID，长度为1到64个字符，只包含字母、数字和\"-\"的字符串。
	UserId string `json:"user_id"`

	// 永久访问密钥ID，即AK。
	AccessKeyId string `json:"access_key_id"`
}

func (o DeleteAccessKeyV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAccessKeyV5Request struct{}"
	}

	return strings.Join([]string{"DeleteAccessKeyV5Request", string(data)}, " ")
}
