package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpIgnoreRulesRequest Request Object
type ShowHttpIgnoreRulesRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	// 规则开关状态
	Status *int32 `json:"status,omitempty"`

	// 分页查询参数，第page页
	Page *int32 `json:"page,omitempty"`

	// 分页查询参数，每页显示pagesize条记录
	Pagesize *int32 `json:"pagesize,omitempty"`

	// 不检测模块类型
	Rule *string `json:"rule,omitempty"`
}

func (o ShowHttpIgnoreRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpIgnoreRulesRequest struct{}"
	}

	return strings.Join([]string{"ShowHttpIgnoreRulesRequest", string(data)}, " ")
}
