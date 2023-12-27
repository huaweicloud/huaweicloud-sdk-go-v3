package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateShrinkageJobResponse Response Object
type CreateShrinkageJobResponse struct {

	// 缩容变更任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateShrinkageJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShrinkageJobResponse struct{}"
	}

	return strings.Join([]string{"CreateShrinkageJobResponse", string(data)}, " ")
}
