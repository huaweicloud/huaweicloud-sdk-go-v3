package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BCS组织监控数据列表查询请求结构
type ListEntityMetricRequestBody struct {

	// 实体类型，[可选值如下 org     # 节点组织 plugin  # 插件](tag:online) 默认为org
	Type string `json:"type" xml:"type"`

	// 具体实体的名称
	EntityName *string `json:"entity_name,omitempty" xml:"entity_name"`
}

func (o ListEntityMetricRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEntityMetricRequestBody struct{}"
	}

	return strings.Join([]string{"ListEntityMetricRequestBody", string(data)}, " ")
}
