package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStsCredentialProviderResponse Response Object
type DeleteStsCredentialProviderResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteStsCredentialProviderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStsCredentialProviderResponse struct{}"
	}

	return strings.Join([]string{"DeleteStsCredentialProviderResponse", string(data)}, " ")
}
