package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPluginAttachableApisRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量，条目数量小于等于0时，自动转换为20，条目数量大于500时，自动转换为500
	Limit *int32 `json:"limit,omitempty"`

	// 插件编号
	PluginId string `json:"plugin_id"`

	// 发布的环境编号
	EnvId *string `json:"env_id,omitempty"`

	// API名称
	ApiName *string `json:"api_name,omitempty"`

	// API编号
	ApiId *string `json:"api_id,omitempty"`

	// 分组编号
	GroupId *string `json:"group_id,omitempty"`

	// 请求方法
	ReqMethod *string `json:"req_method,omitempty"`

	// 请求路径
	ReqUri *string `json:"req_uri,omitempty"`
}

func (o ListPluginAttachableApisRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPluginAttachableApisRequest struct{}"
	}

	return strings.Join([]string{"ListPluginAttachableApisRequest", string(data)}, " ")
}
