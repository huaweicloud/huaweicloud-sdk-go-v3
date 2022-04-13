package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateExecuteJobRequest struct {
	// 集群ID。获取方法，请参见[获取集群ID](https://support.huaweicloud.com/api-mrs/mrs_02_9001.html)。

	ClusterId string `json:"cluster_id"`

	Body *JobExecution `json:"body,omitempty"`
}

func (o CreateExecuteJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateExecuteJobRequest struct{}"
	}

	return strings.Join([]string{"CreateExecuteJobRequest", string(data)}, " ")
}
