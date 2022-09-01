package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateJobRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id" xml:"cluster_id"`

	// 作业名称。
	JobName string `json:"job_name" xml:"job_name"`

	Body *CdmUpdateJobJsonReq `json:"body,omitempty" xml:"body"`
}

func (o UpdateJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateJobRequest struct{}"
	}

	return strings.Join([]string{"UpdateJobRequest", string(data)}, " ")
}
