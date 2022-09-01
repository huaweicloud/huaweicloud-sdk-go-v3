package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateRunRequest struct {

	// 作业ID。
	JobId string `json:"job_id" xml:"job_id"`

	Body *CreateRunRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateRunRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRunRequest struct{}"
	}

	return strings.Join([]string{"CreateRunRequest", string(data)}, " ")
}
