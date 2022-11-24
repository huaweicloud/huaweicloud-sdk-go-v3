package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobFilterDto struct {

	// 作业名称
	JobName *string `json:"job_name,omitempty"`

	// 计算节点标签
	JobNodeLabels *[]string `json:"job_node_labels,omitempty"`
}

func (o JobFilterDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobFilterDto struct{}"
	}

	return strings.Join([]string{"JobFilterDto", string(data)}, " ")
}
