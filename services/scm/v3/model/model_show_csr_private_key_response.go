package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCsrPrivateKeyResponse Response Object
type ShowCsrPrivateKeyResponse struct {

	// 私钥。
	PrivateKey     *string `json:"private_key,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCsrPrivateKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCsrPrivateKeyResponse struct{}"
	}

	return strings.Join([]string{"ShowCsrPrivateKeyResponse", string(data)}, " ")
}
