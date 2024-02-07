package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateInstanceResponse Response Object
type AssociateInstanceResponse struct {
	GlobalEip *UpdateGlobalEip `json:"global_eip,omitempty"`

	// 本次请求的job id
	JobId *string `json:"job_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AssociateInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateInstanceResponse struct{}"
	}

	return strings.Join([]string{"AssociateInstanceResponse", string(data)}, " ")
}
