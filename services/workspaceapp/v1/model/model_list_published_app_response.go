package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPublishedAppResponse Response Object
type ListPublishedAppResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 查发布的应用列表。
	Items          *[]App `json:"items,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPublishedAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublishedAppResponse struct{}"
	}

	return strings.Join([]string{"ListPublishedAppResponse", string(data)}, " ")
}
