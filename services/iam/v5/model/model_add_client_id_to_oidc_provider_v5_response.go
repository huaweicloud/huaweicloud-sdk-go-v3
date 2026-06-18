package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddClientIdToOidcProviderV5Response Response Object
type AddClientIdToOidcProviderV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o AddClientIdToOidcProviderV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddClientIdToOidcProviderV5Response struct{}"
	}

	return strings.Join([]string{"AddClientIdToOidcProviderV5Response", string(data)}, " ")
}
