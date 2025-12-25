package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlertRulesResponse Response Object
type ListAlertRulesResponse struct {

	// 数量
	Count *int64 `json:"count,omitempty"`

	// 模型记录
	Records        *[]AlertRuleItem `json:"records,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListAlertRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertRulesResponse struct{}"
	}

	return strings.Join([]string{"ListAlertRulesResponse", string(data)}, " ")
}
