package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UnusedIamUserAccessKeyDetails struct {

	// 用户访问密钥唯一标识符（ID）。
	AccessKeyId string `json:"access_key_id"`

	// 用户访问密钥的最后访问时间。
	LastAccessed *sdktime.SdkTime `json:"last_accessed,omitempty"`
}

func (o UnusedIamUserAccessKeyDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnusedIamUserAccessKeyDetails struct{}"
	}

	return strings.Join([]string{"UnusedIamUserAccessKeyDetails", string(data)}, " ")
}
