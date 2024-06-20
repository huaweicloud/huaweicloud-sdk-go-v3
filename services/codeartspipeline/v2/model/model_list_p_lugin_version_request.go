package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPLuginVersionRequest Request Object
type ListPLuginVersionRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 插件名
	PluginName string `json:"plugin_name"`

	// 偏移
	Offset string `json:"offset"`

	// 大小
	Limit string `json:"limit"`
}

func (o ListPLuginVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPLuginVersionRequest struct{}"
	}

	return strings.Join([]string{"ListPLuginVersionRequest", string(data)}, " ")
}
