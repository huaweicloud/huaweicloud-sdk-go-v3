package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SearchPublicationsResponse struct {
	// 页面起始页，从0开始

	Offset *int32 `json:"offset,omitempty"`
	// 每页显示的条目数量。 默认值：10。

	Limit *int32 `json:"limit,omitempty"`
	// 总数量。

	Count *int32 `json:"count,omitempty"`
	// 发布响应信息

	Data           *[]PublicationResponseBase `json:"data,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o SearchPublicationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchPublicationsResponse struct{}"
	}

	return strings.Join([]string{"SearchPublicationsResponse", string(data)}, " ")
}
