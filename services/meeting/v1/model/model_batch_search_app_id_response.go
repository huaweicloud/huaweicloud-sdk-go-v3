package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSearchAppIdResponse Response Object
type BatchSearchAppIdResponse struct {

	// 页面起始页，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。 默认值：10。
	Limit *int32 `json:"limit,omitempty"`

	// 总数量。
	Count *int32 `json:"count,omitempty"`

	// 企业应用信息
	Data           *[]AppIdInfoDto `json:"data,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o BatchSearchAppIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSearchAppIdResponse struct{}"
	}

	return strings.Join([]string{"BatchSearchAppIdResponse", string(data)}, " ")
}
