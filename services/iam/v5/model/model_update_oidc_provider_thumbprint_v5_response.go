package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOidcProviderThumbprintV5Response Response Object
type UpdateOidcProviderThumbprintV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateOidcProviderThumbprintV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOidcProviderThumbprintV5Response struct{}"
	}

	return strings.Join([]string{"UpdateOidcProviderThumbprintV5Response", string(data)}, " ")
}
