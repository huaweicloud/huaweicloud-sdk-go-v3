package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePrivateProviderMetadataResponse Response Object
type UpdatePrivateProviderMetadataResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdatePrivateProviderMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateProviderMetadataResponse struct{}"
	}

	return strings.Join([]string{"UpdatePrivateProviderMetadataResponse", string(data)}, " ")
}
