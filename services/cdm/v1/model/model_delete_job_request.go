package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteJobRequest Request Object
type DeleteJobRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 作业名称
	JobName string `json:"job_name"`
}

func (o DeleteJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteJobRequest struct{}"
	}

	return strings.Join([]string{"DeleteJobRequest", string(data)}, " ")
}
