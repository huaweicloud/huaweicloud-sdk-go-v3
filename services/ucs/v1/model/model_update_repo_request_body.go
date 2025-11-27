package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRepoRequestBody struct {

	// 仓库名称
	Name *string `json:"name,omitempty"`

	// 所在命名空间
	Namespace *string `json:"namespace,omitempty"`

	// 仓库类型，包括Bucket、HelmChart、GitRepository、HelmRepository，目前仅支持GitRepository
	RepoType *string `json:"repoType,omitempty"`

	GitRepositorySpec *GitRepositorySpec `json:"gitRepositorySpec,omitempty"`

	SecretInfo *SecretInfo `json:"secretInfo,omitempty"`
}

func (o UpdateRepoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRepoRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateRepoRequestBody", string(data)}, " ")
}
