package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveClientIdFromOidcProviderV5Response Response Object
type RemoveClientIdFromOidcProviderV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o RemoveClientIdFromOidcProviderV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveClientIdFromOidcProviderV5Response struct{}"
	}

	return strings.Join([]string{"RemoveClientIdFromOidcProviderV5Response", string(data)}, " ")
}
