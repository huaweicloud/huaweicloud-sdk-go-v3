package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAddressSetInfoUsingPutResponse Response Object
type UpdateAddressSetInfoUsingPutResponse struct {
	Data           *IdObject `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdateAddressSetInfoUsingPutResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAddressSetInfoUsingPutResponse struct{}"
	}

	return strings.Join([]string{"UpdateAddressSetInfoUsingPutResponse", string(data)}, " ")
}
