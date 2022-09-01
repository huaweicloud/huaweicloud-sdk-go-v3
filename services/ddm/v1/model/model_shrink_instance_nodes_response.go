package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShrinkInstanceNodesResponse struct {

	// DDM实例ID。
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId"`

	// DDM实例名称。
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName"`

	// 任务ID。
	JobId          *string `json:"jobId,omitempty" xml:"jobId"`
	HttpStatusCode int     `json:"-"`
}

func (o ShrinkInstanceNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShrinkInstanceNodesResponse struct{}"
	}

	return strings.Join([]string{"ShrinkInstanceNodesResponse", string(data)}, " ")
}
