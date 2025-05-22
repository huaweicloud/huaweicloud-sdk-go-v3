package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFactoryPendingItemsResponse Response Object
type ListFactoryPendingItemsResponse struct {

	// 待发布包信息
	Data *[]ListFactoryPendingItemsRespData `json:"data,omitempty"`

	// 待发布包数量
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListFactoryPendingItemsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactoryPendingItemsResponse struct{}"
	}

	return strings.Join([]string{"ListFactoryPendingItemsResponse", string(data)}, " ")
}
