package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量作业参数
type BatchJobRequest struct {

	// 批量作业名称
	JobName string `json:"job_name"`

	// 批量作业类型，支持以下选项： - node_upgrade： 节点升级 - deployment_deploy：应用部署 - deployment_upgrade：应用升级
	JobType string `json:"job_type"`

	JobContent *JobContentInfo `json:"job_content,omitempty"`

	// 批量作业描述
	Description *string `json:"description,omitempty"`

	// 批量作业标签
	Tags *[]Attributes `json:"tags,omitempty"`
}

func (o BatchJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchJobRequest struct{}"
	}

	return strings.Join([]string{"BatchJobRequest", string(data)}, " ")
}
