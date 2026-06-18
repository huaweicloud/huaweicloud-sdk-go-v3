package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOidcProviderV5Response Response Object
type DeleteOidcProviderV5Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteOidcProviderV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOidcProviderV5Response struct{}"
	}

	return strings.Join([]string{"DeleteOidcProviderV5Response", string(data)}, " ")
}
