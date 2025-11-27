package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceSummary ResourceSummary 表示成员集群中的资源汇总。
type ResourceSummary struct {

	// 可分配的资源
	Allocatable map[string]interface{} `json:"allocatable,omitempty"`

	// 分配中的资源
	Allocating map[string]interface{} `json:"allocating,omitempty"`

	// 已分配的资源
	Allocated map[string]interface{} `json:"allocated,omitempty"`

	// 资源总量
	Capacity map[string]interface{} `json:"capacity,omitempty"`
}

func (o ResourceSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceSummary struct{}"
	}

	return strings.Join([]string{"ResourceSummary", string(data)}, " ")
}
