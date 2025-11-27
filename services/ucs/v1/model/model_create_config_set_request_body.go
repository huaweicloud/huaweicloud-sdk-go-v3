package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConfigSetRequestBody 创建配置集请求
type CreateConfigSetRequestBody struct {

	// 配置集合的名称
	Name string `json:"name"`

	// 命名空间
	Namespace string `json:"namespace"`

	// 配置集合的类型
	ConfigSetType *string `json:"configSetType,omitempty"`

	// 源代码仓库名称
	RepoName *string `json:"repoName,omitempty"`

	// 对象存储桶的基本信息
	BucketSpec *interface{} `json:"bucketSpec,omitempty"`

	// Helm Chart源基本信息
	HelmChartSpec *interface{} `json:"helmChartSpec,omitempty"`

	GitRepositorySpec *GitRepositorySpec `json:"gitRepositorySpec,omitempty"`

	// Helm仓库基本信息
	HelmRepositorySpec *interface{} `json:"helmRepositorySpec,omitempty"`

	KustomizationSpec *KustomizationSpec `json:"kustomizationSpec,omitempty"`

	ClusterInfo *ClusterInfo `json:"clusterInfo,omitempty"`

	SecretInfo *SecretInfo `json:"secretInfo,omitempty"`
}

func (o CreateConfigSetRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigSetRequestBody struct{}"
	}

	return strings.Join([]string{"CreateConfigSetRequestBody", string(data)}, " ")
}
