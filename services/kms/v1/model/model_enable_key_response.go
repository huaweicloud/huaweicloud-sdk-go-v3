package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type EnableKeyResponse struct {
	KeyInfo        *KeyStatusInfo `json:"key_info,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o EnableKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableKeyResponse struct{}"
	}

	return strings.Join([]string{"EnableKeyResponse", string(data)}, " ")
}
