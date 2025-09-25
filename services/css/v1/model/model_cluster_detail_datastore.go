package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterDetailDatastore 数据搜索引擎类型。
type ClusterDetailDatastore struct {

	// 引擎类型，目前只支持elasticsearch。
	Type *string `json:"type,omitempty"`

	// CSS集群引擎版本号。详细请参考CSS[支持的集群版本](css_03_0056.xml)。
	Version *string `json:"version,omitempty"`

	// 集群发布版本号。
	SubVersion *string `json:"subVersion,omitempty"`

	// 表示集群发布版本是否EOS。
	IsEosCluster *bool `json:"isEosCluster,omitempty"`

	// 集群认证是否支持安全模式。
	SupportSecuritymode *bool `json:"supportSecuritymode,omitempty"`
}

func (o ClusterDetailDatastore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterDetailDatastore struct{}"
	}

	return strings.Join([]string{"ClusterDetailDatastore", string(data)}, " ")
}
