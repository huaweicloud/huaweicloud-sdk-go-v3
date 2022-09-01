package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowKeyStoreResponse struct {
	Keystore       *KeystoreDetails `json:"keystore,omitempty" xml:"keystore"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowKeyStoreResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKeyStoreResponse struct{}"
	}

	return strings.Join([]string{"ShowKeyStoreResponse", string(data)}, " ")
}