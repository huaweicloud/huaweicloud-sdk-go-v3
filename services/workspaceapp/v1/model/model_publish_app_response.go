package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishAppResponse Response Object
type PublishAppResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 查发布的应用列表。
	Items          *[]App `json:"items,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o PublishAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishAppResponse struct{}"
	}

	return strings.Join([]string{"PublishAppResponse", string(data)}, " ")
}
