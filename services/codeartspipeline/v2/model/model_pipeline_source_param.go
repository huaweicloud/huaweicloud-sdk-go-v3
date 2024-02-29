package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipelineSourceParam 流水线源参数
type PipelineSourceParam struct {

	// 代码/制品源参数 - 代码仓/制品源别名。别名仅支持输入大小写英文字母、数字、“_”，至多128个字符
	Alias *string `json:"alias,omitempty"`

	// 代码源参数 - git类型
	GitType *string `json:"git_type,omitempty"`

	// 代码源参数 - Repo代码仓ID
	CodehubId *string `json:"codehub_id,omitempty"`

	// 代码源参数 - 扩展点id
	EndpointId *string `json:"endpoint_id,omitempty"`

	// 代码源参数 - 默认分支
	DefaultBranch *string `json:"default_branch,omitempty"`

	// 代码源参数 - git链接
	GitUrl *string `json:"git_url,omitempty"`

	// 代码源参数 - ssh_git链接
	SshGitUrl *string `json:"ssh_git_url,omitempty"`

	// 代码源参数 - 网页url
	WebUrl *string `json:"web_url,omitempty"`

	// 代码源参数 - 流水线源名称
	RepoName *string `json:"repo_name,omitempty"`

	// 制品源参数 - 制品源类型，generic/docker
	ArtifactType *string `json:"artifact_type,omitempty"`

	// 制品源参数 - 制品源类型名
	ArtifactTypeName *string `json:"artifact_type_name,omitempty"`

	// 制品源参数 - 过滤分支
	BranchFilter *string `json:"branch_filter,omitempty"`

	// 制品源参数 - 目录
	Directory *string `json:"directory,omitempty"`

	// 制品源参数 - 目录ID
	DirectoryId *string `json:"directory_id,omitempty"`

	// 制品源参数 - Docker组织
	Organization *string `json:"organization,omitempty"`

	// 制品源参数 - 包名称
	PackageName *string `json:"package_name,omitempty"`

	// 制品源参数 - 版本
	Version *string `json:"version,omitempty"`

	// 制品源参数 - 获取制品源版本的策略，latest/specificVersion
	VersionStrategy *string `json:"version_strategy,omitempty"`

	// 制品源参数 - 制品源名称,如CloudArtifact
	SourceSystem *string `json:"source_system,omitempty"`
}

func (o PipelineSourceParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineSourceParam struct{}"
	}

	return strings.Join([]string{"PipelineSourceParam", string(data)}, " ")
}
