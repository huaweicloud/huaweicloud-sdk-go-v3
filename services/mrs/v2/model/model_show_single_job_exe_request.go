package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSingleJobExeRequest Request Object
type ShowSingleJobExeRequest struct {

	// 作业ID。获取方法，请参见[获取作业ID](https://support.huaweicloud.com/api-mrs/mrs_02_9001.html)。
	JobExecutionId string `json:"job_execution_id"`

	// 集群ID。获取方法，请参见[获取集群ID](https://support.huaweicloud.com/api-mrs/mrs_02_9001.html)。
	ClusterId string `json:"cluster_id"`
}

func (o ShowSingleJobExeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSingleJobExeRequest struct{}"
	}

	return strings.Join([]string{"ShowSingleJobExeRequest", string(data)}, " ")
}
