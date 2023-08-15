package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneUpdateProtocolRequestBody
type KeystoneUpdateProtocolRequestBody struct {
	Protocol *ProtocolOption `json:"protocol"`
}

func (o KeystoneUpdateProtocolRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneUpdateProtocolRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateProtocolRequestBody", string(data)}, " ")
}
