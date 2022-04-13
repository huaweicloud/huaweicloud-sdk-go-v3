package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type ResetServerPasswordRequestBody struct {
	ResetPassword *ResetServerPasswordOption `json:"reset-password"`
}

func (o ResetServerPasswordRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetServerPasswordRequestBody struct{}"
	}

	return strings.Join([]string{"ResetServerPasswordRequestBody", string(data)}, " ")
}
