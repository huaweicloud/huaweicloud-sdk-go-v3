package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPluginsRequest Request Object
type ListPluginsRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量，条目数量小于等于0时，自动转换为20，条目数量大于500时，自动转换为500
	Limit *int32 `json:"limit,omitempty"`

	// 插件类型
	PluginType *string `json:"plugin_type,omitempty"`

	// 插件可见范围
	PluginScope *string `json:"plugin_scope,omitempty"`

	// 插件编码
	PluginId *string `json:"plugin_id,omitempty"`

	// 插件名称，支持模糊查询
	PluginName *string `json:"plugin_name,omitempty"`

	// 指定需要精确匹配查找的参数名称，目前支持插件名称
	PreciseSearch *string `json:"precise_search,omitempty"`
}

func (o ListPluginsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPluginsRequest struct{}"
	}

	return strings.Join([]string{"ListPluginsRequest", string(data)}, " ")
}
