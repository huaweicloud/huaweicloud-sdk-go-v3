package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOidcProviderV5Response Response Object
type UpdateOidcProviderV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateOidcProviderV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOidcProviderV5Response struct{}"
	}

	return strings.Join([]string{"UpdateOidcProviderV5Response", string(data)}, " ")
}
