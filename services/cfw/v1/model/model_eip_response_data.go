package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EipResponseData struct {

	// 每页显示个数
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	// eip资源记录
	Records *[]EipResource `json:"records,omitempty"`
}

func (o EipResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipResponseData struct{}"
	}

	return strings.Join([]string{"EipResponseData", string(data)}, " ")
}
