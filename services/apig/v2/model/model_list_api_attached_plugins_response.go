package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListApiAttachedPluginsResponse struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 绑定API的插件列表。
	Plugins        *[]AttachedPluginInfo `json:"plugins,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListApiAttachedPluginsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiAttachedPluginsResponse struct{}"
	}

	return strings.Join([]string{"ListApiAttachedPluginsResponse", string(data)}, " ")
}
