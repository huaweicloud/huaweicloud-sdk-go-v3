package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReqUpdateAssociatedResourceRules 更新规则入参
type ReqUpdateAssociatedResourceRules struct {

	// 批量更新的规则信息
	Rules *[]ReqAssociatedResourceRule `json:"rules,omitempty"`
}

func (o ReqUpdateAssociatedResourceRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReqUpdateAssociatedResourceRules struct{}"
	}

	return strings.Join([]string{"ReqUpdateAssociatedResourceRules", string(data)}, " ")
}
