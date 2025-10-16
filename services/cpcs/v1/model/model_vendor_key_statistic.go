package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VendorKeyStatistic struct {

	// 总数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 有效数量
	ValidCount *int32 `json:"valid_count,omitempty"`

	// 无效数量
	InvalidCount *int32 `json:"invalid_count,omitempty"`

	// 租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 服务类型
	ServerType *string `json:"server_type,omitempty"`

	// 算法
	Algorithm *string `json:"algorithm,omitempty"`
}

func (o VendorKeyStatistic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VendorKeyStatistic struct{}"
	}

	return strings.Join([]string{"VendorKeyStatistic", string(data)}, " ")
}
