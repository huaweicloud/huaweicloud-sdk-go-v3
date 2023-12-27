package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBuildJobScm 构建执行SCM
type CreateBuildJobScm struct {

	// 代码分支
	Branch *string `json:"branch,omitempty"`

	// 代码仓地址
	Url string `json:"url"`

	// repo的id
	RepoId *string `json:"repo_id,omitempty"`

	// 代码仓http地址
	WebUrl *string `json:"web_url,omitempty"`

	// 仓库类别，Repo、Github等
	ScmType string `json:"scm_type"`

	// 是否自动构建
	IsAutoBuild *bool `json:"is_auto_build,omitempty"`

	// 构建类别
	BuildType *string `json:"build_type,omitempty"`

	// 克隆深度
	Depth *string `json:"depth,omitempty"`

	// endpointId
	EndPointId *string `json:"end_point_id,omitempty"`

	// source
	Source *string `json:"source,omitempty"`
}

func (o CreateBuildJobScm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBuildJobScm struct{}"
	}

	return strings.Join([]string{"CreateBuildJobScm", string(data)}, " ")
}
