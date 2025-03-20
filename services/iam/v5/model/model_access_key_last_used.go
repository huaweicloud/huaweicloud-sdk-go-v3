package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessKeyLastUsed 访问密钥的最后使用时间。
type AccessKeyLastUsed struct {

	// 访问密钥的最后使用时间。若不存在则表示从未使用过。
	LastUsedAt *sdktime.SdkTime `json:"last_used_at,omitempty"`
}

func (o AccessKeyLastUsed) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessKeyLastUsed struct{}"
	}

	return strings.Join([]string{"AccessKeyLastUsed", string(data)}, " ")
}
