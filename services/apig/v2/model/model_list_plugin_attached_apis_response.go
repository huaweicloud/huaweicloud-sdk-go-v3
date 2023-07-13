package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPluginAttachedApisResponse Response Object
type ListPluginAttachedApisResponse struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 绑定插件的API列表。
	Apis           *[]PluginApiInfo `json:"apis,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListPluginAttachedApisResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPluginAttachedApisResponse struct{}"
	}

	return strings.Join([]string{"ListPluginAttachedApisResponse", string(data)}, " ")
}
