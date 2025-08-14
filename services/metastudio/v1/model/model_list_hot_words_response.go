package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHotWordsResponse Response Object
type ListHotWordsResponse struct {

	// 页面起始页,从0开始
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// 总数量
	Count *int32 `json:"count,omitempty"`

	// 热词记录信息
	Data *[]HotWordsInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListHotWordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHotWordsResponse struct{}"
	}

	return strings.Join([]string{"ListHotWordsResponse", string(data)}, " ")
}
