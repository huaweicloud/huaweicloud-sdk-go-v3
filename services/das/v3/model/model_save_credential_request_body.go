package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SaveCredentialRequestBody struct {

	// AK
	Ak string `json:"ak"`

	// SK
	Sk string `json:"sk"`
}

func (o SaveCredentialRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveCredentialRequestBody struct{}"
	}

	return strings.Join([]string{"SaveCredentialRequestBody", string(data)}, " ")
}
