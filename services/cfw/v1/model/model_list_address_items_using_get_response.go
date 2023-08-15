package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAddressItemsUsingGetResponse Response Object
type ListAddressItemsUsingGetResponse struct {
	Data           *AddressItemListResponseDtoData `json:"data,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListAddressItemsUsingGetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAddressItemsUsingGetResponse struct{}"
	}

	return strings.Join([]string{"ListAddressItemsUsingGetResponse", string(data)}, " ")
}
