package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFacialAnimationsResponse Response Object
type CreateFacialAnimationsResponse struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateFacialAnimationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFacialAnimationsResponse struct{}"
	}

	return strings.Join([]string{"CreateFacialAnimationsResponse", string(data)}, " ")
}
