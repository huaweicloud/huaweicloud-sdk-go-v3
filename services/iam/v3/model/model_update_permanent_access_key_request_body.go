package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type UpdatePermanentAccessKeyRequestBody struct {
	Credential *UpdateCredentialOption `json:"credential"`
}

func (o UpdatePermanentAccessKeyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePermanentAccessKeyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePermanentAccessKeyRequestBody", string(data)}, " ")
}
