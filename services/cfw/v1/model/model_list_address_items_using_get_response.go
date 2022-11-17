package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
