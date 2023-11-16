package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBasePluginsRequest Request Object
type ListBasePluginsRequest struct {

	// 租户ID
	DomainId string `json:"domain_id"`

	// 属性
	Attribution *string `json:"attribution,omitempty"`

	// 偏移
	Offset *string `json:"offset,omitempty"`

	// 大小
	Limit *string `json:"limit,omitempty"`
}

func (o ListBasePluginsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBasePluginsRequest struct{}"
	}

	return strings.Join([]string{"ListBasePluginsRequest", string(data)}, " ")
}
