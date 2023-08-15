package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopJobRequest Request Object
type StopJobRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 作业名称
	JobName string `json:"job_name"`
}

func (o StopJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopJobRequest struct{}"
	}

	return strings.Join([]string{"StopJobRequest", string(data)}, " ")
}
