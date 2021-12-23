package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShrinkInstanceNodeResponse struct {
	// 任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShrinkInstanceNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShrinkInstanceNodeResponse struct{}"
	}

	return strings.Join([]string{"ShrinkInstanceNodeResponse", string(data)}, " ")
}
