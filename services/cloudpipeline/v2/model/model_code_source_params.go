package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 源参数
type CodeSourceParams struct {

	// 代码仓类型
	GitType *string `json:"git_type,omitempty"`

	// codehub代码仓ID
	CodehubId *string `json:"codehub_id,omitempty"`

	// 扩展点ID
	EndpointId *string `json:"endpoint_id,omitempty"`

	// 默认分支
	DefaultBranch *string `json:"default_branch,omitempty"`

	// 代码仓https地址
	GitUrl *string `json:"git_url,omitempty"`

	// 代码仓ssh地址
	SshGitUrl *string `json:"ssh_git_url,omitempty"`

	// 代码仓页面地址
	WebUrl *string `json:"web_url,omitempty"`

	// 代码仓名称
	RepoName *string `json:"repo_name,omitempty"`

	// 代码仓别名
	Alias *string `json:"alias,omitempty"`
}

func (o CodeSourceParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CodeSourceParams struct{}"
	}

	return strings.Join([]string{"CodeSourceParams", string(data)}, " ")
}
