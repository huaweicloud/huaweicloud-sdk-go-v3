package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateTaskRequestV2 struct {

	// 检查类型，数组格式，默认为source
	CheckType *[]string `json:"check_type,omitempty" xml:"check_type"`

	// 仓库地址
	GitUrl string `json:"git_url" xml:"git_url"`

	// 仓库分支
	GitBranch string `json:"git_branch" xml:"git_branch"`

	// 检查语言，数组格式，支持cpp,java,js,python,php,css,html,go,typescript,csharp
	Language []string `json:"language" xml:"language"`

	// 指定规则集
	RuleSets *[]RuleSetV2 `json:"rule_sets,omitempty" xml:"rule_sets"`

	// 检查类型，支持full/inc两种类型，full表示全量检查，inc表示mr检查
	TaskType *string `json:"task_type,omitempty" xml:"task_type"`

	// 仓库有权限的用户名
	Username *string `json:"username,omitempty" xml:"username"`

	// 仓库有权限的用户token
	AccessToken *string `json:"access_token,omitempty" xml:"access_token"`

	// 仓库有权限的用户endpointId
	EndpointId *string `json:"endpoint_id,omitempty" xml:"endpoint_id"`

	IncConfig *IncConfigV2 `json:"inc_config,omitempty" xml:"inc_config"`

	// 是否打开fossbot检查,默认不开
	EnableFossbot *bool `json:"enable_fossbot,omitempty" xml:"enable_fossbot"`

	// 资源池id，可以从资源池管理页面获取
	ResourcePoolId *string `json:"resource_pool_id,omitempty" xml:"resource_pool_id"`
}

func (o CreateTaskRequestV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTaskRequestV2 struct{}"
	}

	return strings.Join([]string{"CreateTaskRequestV2", string(data)}, " ")
}
