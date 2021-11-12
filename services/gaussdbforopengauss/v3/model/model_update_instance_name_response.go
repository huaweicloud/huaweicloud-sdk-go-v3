package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateInstanceNameResponse struct {
	// 修改实例名称的任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateInstanceNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateInstanceNameResponse", string(data)}, " ")
}
