package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WaterfallPageListParam struct {

	// 页面的分页标志位；为分页的最后一条记录的id
	Marker *string `json:"marker,omitempty"`

	// 查询返回记录的数量限制
	Limit *int32 `json:"limit,omitempty"`
}

func (o WaterfallPageListParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WaterfallPageListParam struct{}"
	}

	return strings.Join([]string{"WaterfallPageListParam", string(data)}, " ")
}
