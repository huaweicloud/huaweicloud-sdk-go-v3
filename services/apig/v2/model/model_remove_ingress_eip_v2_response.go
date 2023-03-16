package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RemoveIngressEipV2Response struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 公网入口变更的任务信息
	Message *string `json:"message,omitempty"`

	// 任务编号
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RemoveIngressEipV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveIngressEipV2Response struct{}"
	}

	return strings.Join([]string{"RemoveIngressEipV2Response", string(data)}, " ")
}
