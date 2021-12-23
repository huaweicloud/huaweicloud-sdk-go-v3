package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunInstanceActionResponse struct {
	// 实例删除的任务id。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunInstanceActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunInstanceActionResponse struct{}"
	}

	return strings.Join([]string{"RunInstanceActionResponse", string(data)}, " ")
}
