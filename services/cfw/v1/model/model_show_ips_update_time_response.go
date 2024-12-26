package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIpsUpdateTimeResponse Response Object
type ShowIpsUpdateTimeResponse struct {
	Data *[]IpsRuleUpdateTimeVo `json:"data,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorDescription *string `json:"error_description,omitempty"`

	FailReason *string `json:"fail_reason,omitempty"`

	JobId *string `json:"job_id,omitempty"`

	OrderId *string `json:"order_id,omitempty"`

	TraceId        *string `json:"trace_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowIpsUpdateTimeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpsUpdateTimeResponse struct{}"
	}

	return strings.Join([]string{"ShowIpsUpdateTimeResponse", string(data)}, " ")
}
