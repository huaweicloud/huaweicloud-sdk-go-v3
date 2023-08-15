package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlertRulesResponse Response Object
type ListAlertRulesResponse struct {

	// total count
	Count *int64 `json:"count,omitempty"`

	// rules
	Records *[]AlertRule `json:"records,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAlertRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertRulesResponse struct{}"
	}

	return strings.Join([]string{"ListAlertRulesResponse", string(data)}, " ")
}
