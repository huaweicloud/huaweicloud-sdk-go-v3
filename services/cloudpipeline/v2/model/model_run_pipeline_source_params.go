package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 源参数
type RunPipelineSourceParams struct {

	// 代码仓类型
	GitType *string `json:"git_type,omitempty"`

	// 代码仓https地址
	GitUrl *string `json:"git_url,omitempty"`

	// 代码仓ssh地址
	SshGitUrl *string `json:"ssh_git_url,omitempty"`

	// 代码仓页面地址
	WebUrl *string `json:"web_url,omitempty"`

	// 代码仓名
	RepoName *string `json:"repo_name,omitempty"`

	// 默认分支
	DefaultBranch *string `json:"default_branch,omitempty"`

	// 扩展点ID
	EndpointId *string `json:"endpoint_id,omitempty"`

	// codehub代码仓ID
	CodehubId *string `json:"codehub_id,omitempty"`

	// 代码仓别名
	Alias *string `json:"alias,omitempty"`

	BuildParams *RunPipelineSourceParamsBuildParams `json:"build_params,omitempty"`
}

func (o RunPipelineSourceParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunPipelineSourceParams struct{}"
	}

	return strings.Join([]string{"RunPipelineSourceParams", string(data)}, " ")
}
