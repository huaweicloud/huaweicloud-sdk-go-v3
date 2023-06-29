package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClearGraph2Response Response Object
type ClearGraph2Response struct {

	// 执行该异步任务的jobId。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ClearGraph2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClearGraph2Response struct{}"
	}

	return strings.Join([]string{"ClearGraph2Response", string(data)}, " ")
}
