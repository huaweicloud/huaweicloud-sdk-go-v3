package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFeatureGatesResponse Response Object
type ListFeatureGatesResponse struct {

	// 是否开启域名管理
	EnableDomainName *bool `json:"enableDomainName,omitempty"`

	// 老化策略是否支持多条规则
	EnableCombinationRetention *bool `json:"enableCombinationRetention,omitempty"`

	// 是否支持制品扫描
	EnableArtifactScanning *bool `json:"enableArtifactScanning,omitempty"`

	// 是否支持制品清理
	EnableArtifactGc *bool `json:"enableArtifactGc,omitempty"`

	// 仓库是否支持标签
	EnableRepositoryTag *bool `json:"enableRepositoryTag,omitempty"`

	// 创建命名空间时是否支持传入标签以及查询命名空间列表是否支持根据标签过滤
	EnableNamespaceTag *bool `json:"enableNamespaceTag,omitempty"`

	// 是否支持查询制品的Manifest信息
	EnableArtifactMetadata *bool `json:"enableArtifactMetadata,omitempty"`
	HttpStatusCode         int   `json:"-"`
}

func (o ListFeatureGatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFeatureGatesResponse struct{}"
	}

	return strings.Join([]string{"ListFeatureGatesResponse", string(data)}, " ")
}
