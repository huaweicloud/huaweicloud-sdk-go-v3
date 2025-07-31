package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterProtectionItemResponse Response Object
type ListClusterProtectionItemResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 漏洞信息
	Vuls *[]string `json:"vuls,omitempty"`

	// 基线信息
	Baselines *[]ClusterBaselineResponseInfo `json:"baselines,omitempty"`

	// 恶意文件信息
	Malwares *[]ClusterMalwareResponseInfo `json:"malwares,omitempty"`

	// 镜像信息
	Images *[]ClusterImageResponseInfo `json:"images,omitempty"`

	// 集群信息
	Clusters       *[]ClusterItemResponseInfo `json:"clusters,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListClusterProtectionItemResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterProtectionItemResponse struct{}"
	}

	return strings.Join([]string{"ListClusterProtectionItemResponse", string(data)}, " ")
}
