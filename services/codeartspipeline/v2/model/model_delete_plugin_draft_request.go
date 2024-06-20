package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePluginDraftRequest Request Object
type DeletePluginDraftRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 插件名
	PluginName string `json:"plugin_name"`

	// 版本
	Version string `json:"version"`
}

func (o DeletePluginDraftRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePluginDraftRequest struct{}"
	}

	return strings.Join([]string{"DeletePluginDraftRequest", string(data)}, " ")
}
