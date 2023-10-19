package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAddressSetResponse Response Object
type DeleteAddressSetResponse struct {
	Data           *IdObject `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o DeleteAddressSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAddressSetResponse struct{}"
	}

	return strings.Join([]string{"DeleteAddressSetResponse", string(data)}, " ")
}
