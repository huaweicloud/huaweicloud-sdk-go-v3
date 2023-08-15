package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddIngressEipV2Response Response Object
type AddIngressEipV2Response struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 公网入口变更的任务信息
	Message *string `json:"message,omitempty"`

	// 任务编号
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddIngressEipV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddIngressEipV2Response struct{}"
	}

	return strings.Join([]string{"AddIngressEipV2Response", string(data)}, " ")
}
