package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCsrPrivateKeyRequest Request Object
type ShowCsrPrivateKeyRequest struct {

	// CSR的ID。
	Id string `json:"id"`
}

func (o ShowCsrPrivateKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCsrPrivateKeyRequest struct{}"
	}

	return strings.Join([]string{"ShowCsrPrivateKeyRequest", string(data)}, " ")
}
