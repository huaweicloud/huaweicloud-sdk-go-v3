package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePrivateProviderResponse Response Object
type DeletePrivateProviderResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePrivateProviderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateProviderResponse struct{}"
	}

	return strings.Join([]string{"DeletePrivateProviderResponse", string(data)}, " ")
}
