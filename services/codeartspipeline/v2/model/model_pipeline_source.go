package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipelineSource 源码仓库参数
type PipelineSource struct {

	// 源码仓类型
	ScmType string `json:"scmType"`

	// 是否配置Webhook
	HookFlag string `json:"hookFlag"`

	// 默认分支
	DefaultBranch string `json:"defaultBranch"`

	// webhook配置数据
	Trigger string `json:"trigger"`

	// 代码仓别名
	Alias string `json:"alias"`

	// 代码仓显示名
	DisplayName string `json:"displayName"`

	// 源码仓名称
	RepoName string `json:"repoName"`

	// 代码仓ID
	RepoId string `json:"repoId"`

	// 代码仓所有者
	RepoOwner string `json:"repoOwner"`

	// 代码仓访问地址
	GitUrl string `json:"gitUrl"`

	// 代码仓Web页面地址
	WebUrl string `json:"webUrl"`
}

func (o PipelineSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineSource struct{}"
	}

	return strings.Join([]string{"PipelineSource", string(data)}, " ")
}
