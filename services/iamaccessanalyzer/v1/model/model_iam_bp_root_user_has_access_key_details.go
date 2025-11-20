package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IamBpRootUserHasAccessKeyDetails Root用户有访问密钥分析详细结果。
type IamBpRootUserHasAccessKeyDetails struct {

	// 用户访问密钥唯一标识符（ID）。
	AccessKeyId string `json:"access_key_id"`

	// 用户访问密钥的最后访问时间。
	LastAccessed *sdktime.SdkTime `json:"last_accessed,omitempty"`

	// 用户访问密钥的创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`
}

func (o IamBpRootUserHasAccessKeyDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IamBpRootUserHasAccessKeyDetails struct{}"
	}

	return strings.Join([]string{"IamBpRootUserHasAccessKeyDetails", string(data)}, " ")
}
