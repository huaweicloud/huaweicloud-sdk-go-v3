package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrderResponse Response Object
type ListOrderResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`

	// 列表数据
	Data           *[]OrderItem `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrderResponse struct{}"
	}

	return strings.Join([]string{"ListOrderResponse", string(data)}, " ")
}
