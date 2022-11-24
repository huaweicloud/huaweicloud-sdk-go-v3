package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddAddressItemsUsingPostRequest struct {
	Body *AddAddressItemsInfoDto `json:"body,omitempty"`
}

func (o AddAddressItemsUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAddressItemsUsingPostRequest struct{}"
	}

	return strings.Join([]string{"AddAddressItemsUsingPostRequest", string(data)}, " ")
}
