package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCredentialRequest Request Object
type ShowCredentialRequest struct {
}

func (o ShowCredentialRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCredentialRequest struct{}"
	}

	return strings.Join([]string{"ShowCredentialRequest", string(data)}, " ")
}
