package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteServiceItemsResponse Response Object
type BatchDeleteServiceItemsResponse struct {

	// 批量删除服务组成员id列表
	Data           *[]string `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeleteServiceItemsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteServiceItemsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteServiceItemsResponse", string(data)}, " ")
}
