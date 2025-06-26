package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateLongTermCredentialRequestBody struct {

	// 是否启用
	Enable bool `json:"enable"`
}

func (o UpdateLongTermCredentialRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLongTermCredentialRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateLongTermCredentialRequestBody", string(data)}, " ")
}
