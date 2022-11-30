package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageResourceListParam struct {

	// 页面的分页标志位
	Maker *string `json:"maker,omitempty"`

	// 查询返回记录的数量限制
	Limit *int32 `json:"limit,omitempty"`

	// 关键字模糊搜索
	Keywords map[string]string `json:"keywords,omitempty"`

	// 是否需要返回拓扑树,默认是false。需要：true---性能差，不需要false--性能好
	CiRelationships *bool `json:"ci_relationships,omitempty"`

	// 节点类型，取值：application、sub_application、component、environment
	CiType string `json:"ci_type"`

	// 环境的region信息，若没有值，代表全部
	CiRegion *string `json:"ci_region,omitempty"`

	// 节点id
	CiId string `json:"ci_id"`
}

func (o PageResourceListParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageResourceListParam struct{}"
	}

	return strings.Join([]string{"PageResourceListParam", string(data)}, " ")
}
