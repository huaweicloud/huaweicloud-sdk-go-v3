package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPluginVersionNumberResponse Response Object
type ListPluginVersionNumberResponse struct {

	// 偏移
	Offset *int32 `json:"offset,omitempty"`

	// 大小
	Limit *int32 `json:"limit,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 结果集
	Data           *[]string `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListPluginVersionNumberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPluginVersionNumberResponse struct{}"
	}

	return strings.Join([]string{"ListPluginVersionNumberResponse", string(data)}, " ")
}
