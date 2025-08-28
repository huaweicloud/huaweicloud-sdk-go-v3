package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessKey 创建的永久访问密钥。
type AccessKey struct {

	// IAM用户ID。
	UserId string `json:"user_id"`

	// 创建的永久访问密钥ID，即AK。
	AccessKeyId string `json:"access_key_id"`

	// 访问密钥创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 创建的SK。
	SecretAccessKey string `json:"secret_access_key"`

	Status *AccessKeyStatus `json:"status"`
}

func (o AccessKey) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessKey struct{}"
	}

	return strings.Join([]string{"AccessKey", string(data)}, " ")
}
