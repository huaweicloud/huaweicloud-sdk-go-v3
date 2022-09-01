package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowJobStatusRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id" xml:"cluster_id"`

	// 作业名称
	JobName string `json:"job_name" xml:"job_name"`
}

func (o ShowJobStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowJobStatusRequest", string(data)}, " ")
}
