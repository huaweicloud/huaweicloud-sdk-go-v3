package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGlobalEipResponse Response Object
type UpdateGlobalEipResponse struct {
	GlobalEip *UpdateGlobalEip `json:"global_eip,omitempty"`

	// 本次请求的job id
	JobId *string `json:"job_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateGlobalEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalEipResponse struct{}"
	}

	return strings.Join([]string{"UpdateGlobalEipResponse", string(data)}, " ")
}
