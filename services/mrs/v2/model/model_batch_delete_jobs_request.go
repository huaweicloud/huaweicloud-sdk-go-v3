package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDeleteJobsRequest struct {

	// 集群ID。获取方法，请参见[获取集群ID](https://support.huaweicloud.com/api-mrs/mrs_02_9001.html)。
	ClusterId string `json:"cluster_id" xml:"cluster_id"`

	Body *JobBatchDelete `json:"body,omitempty" xml:"body"`
}

func (o BatchDeleteJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteJobsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteJobsRequest", string(data)}, " ")
}
