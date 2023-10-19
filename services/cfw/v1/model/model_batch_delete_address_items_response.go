package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAddressItemsResponse Response Object
type BatchDeleteAddressItemsResponse struct {

	// 批量删除地址组成员id列表
	Data           *[]string `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeleteAddressItemsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAddressItemsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteAddressItemsResponse", string(data)}, " ")
}
