package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListApiAttachablePluginsResponse struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 插件列表。
	Plugins        *[]PluginInfo `json:"plugins,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListApiAttachablePluginsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiAttachablePluginsResponse struct{}"
	}

	return strings.Join([]string{"ListApiAttachablePluginsResponse", string(data)}, " ")
}
