package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRepoRequestBody struct {

	// 仓库名称
	Name *string `json:"name,omitempty"`

	// 所在命名空间
	Namespace *string `json:"namespace,omitempty"`

	GitRepositorySpec *GitRepositorySpec `json:"gitRepositorySpec,omitempty"`

	SecretInfo *SecretInfo `json:"secretInfo,omitempty"`

	ClusterInfo *ClusterInfo `json:"clusterInfo,omitempty"`
}

func (o CreateRepoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRepoRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRepoRequestBody", string(data)}, " ")
}
