package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStatisticSecretKeyResponse Response Object
type ShowStatisticSecretKeyResponse struct {

	// 总数量
	TotalCount *int32 `json:"total_count,omitempty"`

	// 无效数量
	InvalidCount *int32 `json:"invalid_count,omitempty"`

	// 有效数量
	ValidCount *int32 `json:"valid_count,omitempty"`

	// 密钥分布按服务类型统计
	KeyCountsByServiceType *[]VendorKeyStatistic `json:"key_counts_by_service_type,omitempty"`

	// 密钥分布按算法统计
	KeyCountsByAlgorithm *[]VendorKeyStatistic `json:"key_counts_by_algorithm,omitempty"`

	// 密钥分布按算法和集群统计
	KeyCountsByAlgorithmAndCluster *[]VendorKeyStatistic `json:"key_counts_by_algorithm_and_cluster,omitempty"`
	HttpStatusCode                 int                   `json:"-"`
}

func (o ShowStatisticSecretKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatisticSecretKeyResponse struct{}"
	}

	return strings.Join([]string{"ShowStatisticSecretKeyResponse", string(data)}, " ")
}
