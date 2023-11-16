package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OpenSourceRuleSetCreateVo 项目级级开源治理策略实例详情
type OpenSourceRuleSetCreateVo struct {

	// 开源治理策略名称
	Name string `json:"name"`

	// 开源治理策略父策略ID
	ParentId *string `json:"parent_id,omitempty"`

	Content *OpenSourceRuleContent `json:"content"`
}

func (o OpenSourceRuleSetCreateVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenSourceRuleSetCreateVo struct{}"
	}

	return strings.Join([]string{"OpenSourceRuleSetCreateVo", string(data)}, " ")
}
