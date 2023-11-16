package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TenantOpenSourceRuleSetCreateVo 租户级开源治理策略实例详情
type TenantOpenSourceRuleSetCreateVo struct {

	// 开源治理策略名称
	Name string `json:"name"`

	Content *OpenSourceRuleContent `json:"content"`
}

func (o TenantOpenSourceRuleSetCreateVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TenantOpenSourceRuleSetCreateVo struct{}"
	}

	return strings.Join([]string{"TenantOpenSourceRuleSetCreateVo", string(data)}, " ")
}
