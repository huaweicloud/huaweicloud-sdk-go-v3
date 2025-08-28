package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessKeyMetadata 永久访问密钥。
type AccessKeyMetadata struct {

	// IAM用户ID。
	UserId string `json:"user_id"`

	// 永久访问密钥ID，即AK。
	AccessKeyId string `json:"access_key_id"`

	// 访问密钥创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	Status *AccessKeyStatus `json:"status"`
}

func (o AccessKeyMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessKeyMetadata struct{}"
	}

	return strings.Join([]string{"AccessKeyMetadata", string(data)}, " ")
}
