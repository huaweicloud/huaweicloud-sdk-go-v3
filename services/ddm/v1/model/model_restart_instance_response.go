package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartInstanceResponse Response Object
type RestartInstanceResponse struct {

	// DDM实例ID。
	InstanceId *string `json:"instanceId,omitempty"`

	// DDM实例名称。
	InstanceName *string `json:"instanceName,omitempty"`

	// 任务ID。
	JobId          *string `json:"jobId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestartInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartInstanceResponse struct{}"
	}

	return strings.Join([]string{"RestartInstanceResponse", string(data)}, " ")
}
