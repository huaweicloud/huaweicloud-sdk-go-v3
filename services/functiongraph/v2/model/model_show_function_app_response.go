package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFunctionAppResponse Response Object
type ShowFunctionAppResponse struct {

	// 应用名称
	Name *string `json:"name,omitempty"`

	// 最后修改时间
	LastModifiedTime *int64 `json:"last_modified_time,omitempty"`

	StackResources *StackResource `json:"stack_resources,omitempty"`

	// 应用状态
	Status *string `json:"status,omitempty"`

	// 资源栈名称
	StackName *string `json:"stack_name,omitempty"`

	// 资源栈id
	StackId *string `json:"stack_id,omitempty"`

	// 存储库名称
	RepoName *string `json:"repo_name,omitempty"`

	// 应用描述
	Description *string `json:"description,omitempty"`

	Repo *RepoInfo `json:"repo,omitempty"`

	// 管道id
	PipelineId *string `json:"pipeline_id,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 调用URL
	ApigUrl        *string `json:"apig_url,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowFunctionAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFunctionAppResponse struct{}"
	}

	return strings.Join([]string{"ShowFunctionAppResponse", string(data)}, " ")
}
