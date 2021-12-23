package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DisableKeyResponse struct {
	KeyInfo        *KeyStatusInfo `json:"key_info,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o DisableKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableKeyResponse struct{}"
	}

	return strings.Join([]string{"DisableKeyResponse", string(data)}, " ")
}
