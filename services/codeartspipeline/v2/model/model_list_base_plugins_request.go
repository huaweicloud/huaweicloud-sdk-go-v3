package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBasePluginsRequest Request Object
type ListBasePluginsRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 插件属性,可选official、published
	Attribution string `json:"attribution"`

	// 偏移
	Offset string `json:"offset"`

	// 大小
	Limit string `json:"limit"`
}

func (o ListBasePluginsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBasePluginsRequest struct{}"
	}

	return strings.Join([]string{"ListBasePluginsRequest", string(data)}, " ")
}
