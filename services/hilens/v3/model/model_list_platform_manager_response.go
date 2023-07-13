package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPlatformManagerResponse Response Object
type ListPlatformManagerResponse struct {

	// 订单总数
	Total *int32 `json:"total,omitempty"`

	// 订单列表
	Data           *[]OrderForm `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListPlatformManagerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlatformManagerResponse struct{}"
	}

	return strings.Join([]string{"ListPlatformManagerResponse", string(data)}, " ")
}
