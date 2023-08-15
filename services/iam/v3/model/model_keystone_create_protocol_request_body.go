package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneCreateProtocolRequestBody
type KeystoneCreateProtocolRequestBody struct {
	Protocol *ProtocolOption `json:"protocol"`
}

func (o KeystoneCreateProtocolRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneCreateProtocolRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneCreateProtocolRequestBody", string(data)}, " ")
}
