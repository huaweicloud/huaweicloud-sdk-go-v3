package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ResizeInstanceVolumeResponse struct {

	// 工作流ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResizeInstanceVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceVolumeResponse struct{}"
	}

	return strings.Join([]string{"ResizeInstanceVolumeResponse", string(data)}, " ")
}
