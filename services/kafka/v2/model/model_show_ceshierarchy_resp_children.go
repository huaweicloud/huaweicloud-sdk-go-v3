package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 子维度信息。
type ShowCeshierarchyRespChildren struct {

	// 子维度名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 监控指标名称列表。
	Metrics *[]string `json:"metrics,omitempty" xml:"metrics"`

	// 监控查询使用的key。
	KeyName *[]string `json:"key_name,omitempty" xml:"key_name"`

	// 监控维度路由。
	DimRouter *[]string `json:"dim_router,omitempty" xml:"dim_router"`
}

func (o ShowCeshierarchyRespChildren) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCeshierarchyRespChildren struct{}"
	}

	return strings.Join([]string{"ShowCeshierarchyRespChildren", string(data)}, " ")
}
