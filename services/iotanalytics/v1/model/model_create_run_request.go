package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRunRequest Request Object
type CreateRunRequest struct {

	// 作业ID。
	JobId string `json:"job_id"`

	Body *CreateRunRequestBody `json:"body,omitempty"`
}

func (o CreateRunRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRunRequest struct{}"
	}

	return strings.Join([]string{"CreateRunRequest", string(data)}, " ")
}
