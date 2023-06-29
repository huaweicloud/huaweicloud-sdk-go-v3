package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEncryptdatasResponse Response Object
type ShowEncryptdatasResponse struct {
	EncryptData    *EncryptData `json:"encrypt_data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowEncryptdatasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEncryptdatasResponse struct{}"
	}

	return strings.Join([]string{"ShowEncryptdatasResponse", string(data)}, " ")
}
