package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateJobRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id" xml:"cluster_id"`

	Body *CdmCreateJobJsonReq `json:"body,omitempty" xml:"body"`
}

func (o CreateJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateJobRequest struct{}"
	}

	return strings.Join([]string{"CreateJobRequest", string(data)}, " ")
}
