package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAddressItemsUsingPostResponse Response Object
type AddAddressItemsUsingPostResponse struct {
	Data           *AddressItems `json:"data,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o AddAddressItemsUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAddressItemsUsingPostResponse struct{}"
	}

	return strings.Join([]string{"AddAddressItemsUsingPostResponse", string(data)}, " ")
}
