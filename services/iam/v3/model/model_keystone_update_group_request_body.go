package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneUpdateGroupRequestBody
type KeystoneUpdateGroupRequestBody struct {
	Group *KeystoneUpdateGroupOption `json:"group"`
}

func (o KeystoneUpdateGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneUpdateGroupRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateGroupRequestBody", string(data)}, " ")
}
