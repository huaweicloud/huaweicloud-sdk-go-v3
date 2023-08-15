package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunRequestV2 代码仓相关的信息非必填，主要是用于CodeFever临时仓库使用，不填时直接使用创建任务时已经提供的信息
type RunRequestV2 struct {

	// 该任务对应临时仓库有权限的用户名
	Username *string `json:"username,omitempty"`

	// 该任务对应临时仓库有权限的用户token
	AccessToken *string `json:"access_token,omitempty"`

	// 该任务对应的临时仓库地址
	GitUrl *string `json:"git_url,omitempty"`

	// 该任务对应的临时仓库分支
	GitBranch *string `json:"git_branch,omitempty"`
}

func (o RunRequestV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunRequestV2 struct{}"
	}

	return strings.Join([]string{"RunRequestV2", string(data)}, " ")
}
