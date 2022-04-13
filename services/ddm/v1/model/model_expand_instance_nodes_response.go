package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExpandInstanceNodesResponse struct {
	// DDM实例ID。

	InstanceId *string `json:"instanceId,omitempty"`
	// DDM实例名称。

	InstanceName *string `json:"instanceName,omitempty"`
	// 任务ID。

	JobId          *string `json:"jobId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExpandInstanceNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandInstanceNodesResponse struct{}"
	}

	return strings.Join([]string{"ExpandInstanceNodesResponse", string(data)}, " ")
}
