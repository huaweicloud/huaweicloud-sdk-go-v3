package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BareMetalModifyPortRequestBody struct {
	Nic *BareMetalModifyPortRequest `json:"nic,omitempty"`
}

func (o BareMetalModifyPortRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BareMetalModifyPortRequestBody struct{}"
	}

	return strings.Join([]string{"BareMetalModifyPortRequestBody", string(data)}, " ")
}
