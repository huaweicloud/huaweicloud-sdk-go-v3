package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepoResponse struct {

	// 仓库的唯一标识符
	Uid *string `json:"uid,omitempty"`

	// 仓库名称
	Name *string `json:"name,omitempty"`

	// 仓库类型，包括Bucket、HelmChart、GitRepository、HelmRepository，目前仅支持GitRepository
	RepoType *string `json:"repoType,omitempty"`

	// 命名空间
	Namespace *string `json:"namespace,omitempty"`

	GitRepository *GitRepository `json:"gitRepository,omitempty"`

	// 仓库状态，包括Health、Failed、Unknown、Progressing
	RepoStatus *string `json:"repoStatus,omitempty"`

	ClusterInfo *ClusterInfo `json:"clusterInfo,omitempty"`

	SecretInfo *SecretInfo `json:"secretInfo,omitempty"`
}

func (o RepoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoResponse struct{}"
	}

	return strings.Join([]string{"RepoResponse", string(data)}, " ")
}
