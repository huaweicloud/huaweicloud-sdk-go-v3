package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImportWorkflowReq struct {

	// 源项目id
	SourceProjectId string `json:"source_project_id"`

	// 源流程id
	SourceWorkflowId string `json:"source_workflow_id"`

	// 目标流程名称 取值范围[1,56]，允许大小写字母、数字、以及特殊字符中划线(-)和下划线(_)。
	DestinationWorkflowName string `json:"destination_workflow_name"`

	// 目标流程版本 取值范围[1,24]，以小写字母或数字或大写字母开头，允许出现中划线，必须以小写字母或数字或大写字母结尾。
	DestinationWorkflowVersion string `json:"destination_workflow_version"`
}

func (o ImportWorkflowReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportWorkflowReq struct{}"
	}

	return strings.Join([]string{"ImportWorkflowReq", string(data)}, " ")
}
