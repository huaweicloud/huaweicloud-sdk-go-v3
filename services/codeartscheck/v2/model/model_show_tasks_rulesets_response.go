package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTasksRulesetsResponse Response Object
type ShowTasksRulesetsResponse struct {

	// 规则集信息
	Info *[]TaskRulesetInfo `json:"info,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowTasksRulesetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTasksRulesetsResponse struct{}"
	}

	return strings.Join([]string{"ShowTasksRulesetsResponse", string(data)}, " ")
}
