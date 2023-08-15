package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateJobInfo struct {

	// 应用名称
	ApplicationName *string `json:"application_name,omitempty"`

	// 任务依赖的模板id
	TemplateId string `json:"template_id"`

	// 应用名称
	ProjectName *string `json:"project_name,omitempty"`

	// 应用代码生成后的地址类型，目前支持0：codehub地址和1：压缩包下载地址
	RepoType *int32 `json:"repo_type,omitempty"`

	// 应用的动态参数json
	Properties *interface{} `json:"properties,omitempty"`

	// 模板 dependency ID 集合
	TemplateDependencies *[]string `json:"template_dependencies,omitempty"`

	RepoInfo *RepositoryInfo `json:"repo_info,omitempty"`
}

func (o TemplateJobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateJobInfo struct{}"
	}

	return strings.Join([]string{"TemplateJobInfo", string(data)}, " ")
}
