package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type PushTranscriberJobsResponse struct {

	// 创建的任务标识, 如果创建任务成功时必须存在。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PushTranscriberJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PushTranscriberJobsResponse struct{}"
	}

	return strings.Join([]string{"PushTranscriberJobsResponse", string(data)}, " ")
}
