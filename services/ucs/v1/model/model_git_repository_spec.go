package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GitRepositorySpec struct {

	// Git仓库地址
	Url *string `json:"url,omitempty"`

	Ref *GitRepositoryRef `json:"ref,omitempty"`

	SecretRef *LocalObjectReference `json:"secretRef,omitempty"`

	// 周期性检查仓库更新的时间间隔，格式如 \"1m\" 表示1分钟
	Interval *string `json:"interval,omitempty"`

	// Git 操作（如克隆）的超时时间，默认60秒
	Timeout *string `json:"timeout,omitempty"`
}

func (o GitRepositorySpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GitRepositorySpec struct{}"
	}

	return strings.Join([]string{"GitRepositorySpec", string(data)}, " ")
}
