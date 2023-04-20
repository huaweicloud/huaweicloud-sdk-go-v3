package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAppInstanceHistoryResponse struct {

	// 应用实例历史版本列表
	History        *[]QueryAppInstanceHistoryResponseDto `json:"history,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o ListAppInstanceHistoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppInstanceHistoryResponse struct{}"
	}

	return strings.Join([]string{"ListAppInstanceHistoryResponse", string(data)}, " ")
}
