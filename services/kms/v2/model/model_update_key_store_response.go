package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKeyStoreResponse Response Object
type UpdateKeyStoreResponse struct {
	Keystore       *KeyStoreStateInfo `json:"keystore,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o UpdateKeyStoreResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeyStoreResponse struct{}"
	}

	return strings.Join([]string{"UpdateKeyStoreResponse", string(data)}, " ")
}
