package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneUpdateProjectRequestBody
type KeystoneUpdateProjectRequestBody struct {
	Project *KeystoneUpdateProjectOption `json:"project"`
}

func (o KeystoneUpdateProjectRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneUpdateProjectRequestBody struct{}"
	}

	return strings.Join([]string{"KeystoneUpdateProjectRequestBody", string(data)}, " ")
}
