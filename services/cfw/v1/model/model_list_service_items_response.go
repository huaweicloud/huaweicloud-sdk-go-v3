package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServiceItemsResponse Response Object
type ListServiceItemsResponse struct {
	Data           *ServiceItemListResponseDtoData `json:"data,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListServiceItemsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceItemsResponse struct{}"
	}

	return strings.Join([]string{"ListServiceItemsResponse", string(data)}, " ")
}
