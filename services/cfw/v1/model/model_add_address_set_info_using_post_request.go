package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddAddressSetInfoUsingPostRequest struct {
	Body *AddAddressSetDto `json:"body,omitempty"`
}

func (o AddAddressSetInfoUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAddressSetInfoUsingPostRequest struct{}"
	}

	return strings.Join([]string{"AddAddressSetInfoUsingPostRequest", string(data)}, " ")
}
