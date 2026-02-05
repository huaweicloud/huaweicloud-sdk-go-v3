package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVoicesRequest Request Object
type ListVoicesRequest struct {

	// 查询已注册的声音列表，每页查询显示的条目数量，默认：10
	Limit int32 `json:"limit"`

	// 查询已注册的声音列表，页码偏移量，表示从此页码偏移量开始查询，offset大于等于0， 默认：0
	Offset string `json:"offset"`
}

func (o ListVoicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVoicesRequest struct{}"
	}

	return strings.Join([]string{"ListVoicesRequest", string(data)}, " ")
}
