package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAddressItemResponse Response Object
type AddAddressItemResponse struct {
	Data           *AddressItems `json:"data,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o AddAddressItemResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAddressItemResponse struct{}"
	}

	return strings.Join([]string{"AddAddressItemResponse", string(data)}, " ")
}
