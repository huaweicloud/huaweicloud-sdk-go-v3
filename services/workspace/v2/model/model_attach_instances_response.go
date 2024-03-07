package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachInstancesResponse Response Object
type AttachInstancesResponse struct {

	// 创建分配桌面总任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AttachInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachInstancesResponse struct{}"
	}

	return strings.Join([]string{"AttachInstancesResponse", string(data)}, " ")
}
