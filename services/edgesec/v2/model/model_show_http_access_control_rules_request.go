package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpAccessControlRulesRequest Request Object
type ShowHttpAccessControlRulesRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 分页查询参数，第page页
	Page *int32 `json:"page,omitempty"`

	// 分页查询参数，每页显示pagesize条记录
	Pagesize *int32 `json:"pagesize,omitempty"`
}

func (o ShowHttpAccessControlRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpAccessControlRulesRequest struct{}"
	}

	return strings.Join([]string{"ShowHttpAccessControlRulesRequest", string(data)}, " ")
}
