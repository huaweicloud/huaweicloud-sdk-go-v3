package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAddressItemsResponse Response Object
type ListAddressItemsResponse struct {
	Data           *AddressItemListResponseDtoData `json:"data,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListAddressItemsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAddressItemsResponse struct{}"
	}

	return strings.Join([]string{"ListAddressItemsResponse", string(data)}, " ")
}
