package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAddressSetResponseData struct {

	// Id
	Id *string `json:"id,omitempty"`
}

func (o UpdateAddressSetResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAddressSetResponseData struct{}"
	}

	return strings.Join([]string{"UpdateAddressSetResponseData", string(data)}, " ")
}
