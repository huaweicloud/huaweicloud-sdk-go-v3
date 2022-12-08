package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteRuleGroupRequest struct {

	// 规则组ID
	GroupId string `json:"group_id"`
}

func (o DeleteRuleGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRuleGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteRuleGroupRequest", string(data)}, " ")
}
