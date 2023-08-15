package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEncryptdatasResponse Response Object
type UpdateEncryptdatasResponse struct {
	EncryptData    *EncryptData `json:"encrypt_data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateEncryptdatasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEncryptdatasResponse struct{}"
	}

	return strings.Join([]string{"UpdateEncryptdatasResponse", string(data)}, " ")
}
