package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExchangesResponse Response Object
type ListExchangesResponse struct {

	// 当前显示数量
	Size *int32 `json:"size,omitempty"`

	// 查询结果总数
	Total *int32 `json:"total,omitempty"`

	// Exchange信息列表
	Items          *[]ExchangeDetails `json:"items,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListExchangesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExchangesResponse struct{}"
	}

	return strings.Join([]string{"ListExchangesResponse", string(data)}, " ")
}
