package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAddressSetInfoUsingPostResponse Response Object
type AddAddressSetInfoUsingPostResponse struct {
	Data           *IdObject `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o AddAddressSetInfoUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAddressSetInfoUsingPostResponse struct{}"
	}

	return strings.Join([]string{"AddAddressSetInfoUsingPostResponse", string(data)}, " ")
}
