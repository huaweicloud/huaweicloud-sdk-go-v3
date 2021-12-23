package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 通过模板创建部署任务请求体
type TemplateTaskRequestBody struct {
	// devcloud创建的项目id

	ProjectId string `json:"project_id"`
	// devcloud创建的项目名称

	ProjectName string `json:"project_name"`
	// 部署模板id

	TemplateId string `json:"template_id"`
	// 部署任务名称

	TaskName string `json:"task_name"`
	// 自定义slave资源池id

	SlaveClusterId *string `json:"slave_cluster_id,omitempty"`
	// 部署参数类

	Configs *[]ConfigInfoDo `json:"configs,omitempty"`
}

func (o TemplateTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateTaskRequestBody struct{}"
	}

	return strings.Join([]string{"TemplateTaskRequestBody", string(data)}, " ")
}
