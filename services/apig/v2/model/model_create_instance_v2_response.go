package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateInstanceV2Response struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id"`

	// 创建实例任务信息
	Message *string `json:"message,omitempty" xml:"message"`

	// 任务编号
	JobId          *string `json:"job_id,omitempty" xml:"job_id"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInstanceV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceV2Response struct{}"
	}

	return strings.Join([]string{"CreateInstanceV2Response", string(data)}, " ")
}
