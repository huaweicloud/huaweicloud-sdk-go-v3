package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAttachInstancesResponse Response Object
type BatchAttachInstancesResponse struct {

	// 分配桌面总任务id
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchAttachInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAttachInstancesResponse struct{}"
	}

	return strings.Join([]string{"BatchAttachInstancesResponse", string(data)}, " ")
}
