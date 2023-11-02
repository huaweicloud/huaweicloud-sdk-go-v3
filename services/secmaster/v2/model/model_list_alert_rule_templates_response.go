package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlertRuleTemplatesResponse Response Object
type ListAlertRuleTemplatesResponse struct {

	// 总数量。Total count.
	Count *int64 `json:"count,omitempty"`

	// 告警规则模板。Alert rule templates.
	Records *[]AlertRuleTemplate `json:"records,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAlertRuleTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertRuleTemplatesResponse struct{}"
	}

	return strings.Join([]string{"ListAlertRuleTemplatesResponse", string(data)}, " ")
}
