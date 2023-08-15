package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServiceItemsDetailsResponse Response Object
type ListServiceItemsDetailsResponse struct {
	Data           *ServiceItemListResponseDtoData `json:"data,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListServiceItemsDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceItemsDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListServiceItemsDetailsResponse", string(data)}, " ")
}
