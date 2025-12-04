package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomerBatchDeleteInstanceRecord struct {

	// 实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 工作流id
	JobId *string `json:"job_id,omitempty"`
}

func (o CustomerBatchDeleteInstanceRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomerBatchDeleteInstanceRecord struct{}"
	}

	return strings.Join([]string{"CustomerBatchDeleteInstanceRecord", string(data)}, " ")
}
