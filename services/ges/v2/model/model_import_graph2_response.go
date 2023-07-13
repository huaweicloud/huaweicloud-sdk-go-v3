package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportGraph2Response Response Object
type ImportGraph2Response struct {

	// 执行该异步任务的jobId。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ImportGraph2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportGraph2Response struct{}"
	}

	return strings.Join([]string{"ImportGraph2Response", string(data)}, " ")
}
