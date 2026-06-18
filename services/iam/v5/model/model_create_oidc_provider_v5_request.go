package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOidcProviderV5Request Request Object
type CreateOidcProviderV5Request struct {
	Body *CreateOidcProviderReqBody `json:"body,omitempty"`
}

func (o CreateOidcProviderV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOidcProviderV5Request struct{}"
	}

	return strings.Join([]string{"CreateOidcProviderV5Request", string(data)}, " ")
}
