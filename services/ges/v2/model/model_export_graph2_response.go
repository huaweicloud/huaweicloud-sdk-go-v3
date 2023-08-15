package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportGraph2Response Response Object
type ExportGraph2Response struct {

	// 执行该异步任务的jobId。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportGraph2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportGraph2Response struct{}"
	}

	return strings.Join([]string{"ExportGraph2Response", string(data)}, " ")
}
