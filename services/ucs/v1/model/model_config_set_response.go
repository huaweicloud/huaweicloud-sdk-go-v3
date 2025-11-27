package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigSetResponse struct {

	// 配置集合的唯一标识
	Uid *string `json:"uid,omitempty"`

	// 配置集合的名称
	Name *string `json:"name,omitempty"`

	// 命名空间
	Namespace *string `json:"namespace,omitempty"`

	// 配置集合的类型
	ConfigSetType *string `json:"configSetType,omitempty"`

	// 仓库名称
	RepoName *string `json:"repoName,omitempty"`

	// bucket基本信息
	Bucket *interface{} `json:"bucket,omitempty"`

	// Helm Chart源基本信息
	HelmChart *interface{} `json:"helmChart,omitempty"`

	GitRepository *GitRepository `json:"gitRepository,omitempty"`

	// Helm仓库的定义与状态等信息
	HelmRepository *interface{} `json:"helmRepository,omitempty"`

	// 仓库状态信息
	RepoStatus *string `json:"repoStatus,omitempty"`

	// Helm Chart的发布配置和状态信息
	HelmRelease *interface{} `json:"helmRelease,omitempty"`

	Kustomization *Kustomization `json:"kustomization,omitempty"`

	// 配置集合状态信息
	ConfigSetStatus *string `json:"configSetStatus,omitempty"`

	ClusterInfo *ClusterInfo `json:"clusterInfo,omitempty"`

	SecretInfo *SecretInfo `json:"secretInfo,omitempty"`
}

func (o ConfigSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigSetResponse struct{}"
	}

	return strings.Join([]string{"ConfigSetResponse", string(data)}, " ")
}
