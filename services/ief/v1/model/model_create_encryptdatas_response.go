package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEncryptdatasResponse Response Object
type CreateEncryptdatasResponse struct {
	EncryptData    *EncryptData `json:"encrypt_data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o CreateEncryptdatasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEncryptdatasResponse struct{}"
	}

	return strings.Join([]string{"CreateEncryptdatasResponse", string(data)}, " ")
}
