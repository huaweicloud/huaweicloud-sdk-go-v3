package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainStatusResult 租户状态信息
type ShowDomainStatusResult struct {

	// 租户状态0:未开通; 1:正常; 2:冻结; 3:关闭; 4、5:注销
	Status float32 `json:"status,omitempty"`

	// 是否包含免费额度
	FreeQuota *bool `json:"free_quota,omitempty"`

	// 是否启用自定义环境
	AllowCustomEnv float32 `json:"allow_custom_env,omitempty"`
}

func (o ShowDomainStatusResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainStatusResult struct{}"
	}

	return strings.Join([]string{"ShowDomainStatusResult", string(data)}, " ")
}
