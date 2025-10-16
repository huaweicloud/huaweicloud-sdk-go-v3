package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowStatisticResourceResponseBodyDatapoints struct {

	// 数据
	Data *int32 `json:"data,omitempty"`

	// 数量
	Count *int64 `json:"count,omitempty"`

	// 接口名称
	ApiName *string `json:"api_name,omitempty"`

	// 租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// 服务类型
	ServerType *string `json:"server_type,omitempty"`
}

func (o ShowStatisticResourceResponseBodyDatapoints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatisticResourceResponseBodyDatapoints struct{}"
	}

	return strings.Join([]string{"ShowStatisticResourceResponseBodyDatapoints", string(data)}, " ")
}
