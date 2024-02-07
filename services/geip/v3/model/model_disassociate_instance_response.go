package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateInstanceResponse Response Object
type DisassociateInstanceResponse struct {
	GlobalEip *UpdateGlobalEip `json:"global_eip,omitempty"`

	// 本次请求的job id
	JobId *string `json:"job_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DisassociateInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateInstanceResponse struct{}"
	}

	return strings.Join([]string{"DisassociateInstanceResponse", string(data)}, " ")
}
