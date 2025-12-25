package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlertRuleTemplatesResponse Response Object
type ListAlertRuleTemplatesResponse struct {

	// 计数
	Count *int64 `json:"count,omitempty"`

	// 记录
	Records        *[]AlertRuleTemplate `json:"records,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListAlertRuleTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertRuleTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListAlertRuleTemplatesResponse", string(data)}, " ")
}
