package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGlobalEipResponse Response Object
type CreateGlobalEipResponse struct {

	// 本次请求的job id
	JobId *string `json:"job_id,omitempty"`

	GlobalEip *CreateGlobalEip `json:"global_eip,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateGlobalEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalEipResponse struct{}"
	}

	return strings.Join([]string{"CreateGlobalEipResponse", string(data)}, " ")
}
