package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBasePluginsResponse Response Object
type ListBasePluginsResponse struct {

	// 偏移
	Offset *int32 `json:"offset,omitempty"`

	// 大小
	Limit *int32 `json:"limit,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 结果集
	Data           *[]PageInfoOptionalSinglePluginVoData `json:"data,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o ListBasePluginsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBasePluginsResponse struct{}"
	}

	return strings.Join([]string{"ListBasePluginsResponse", string(data)}, " ")
}
