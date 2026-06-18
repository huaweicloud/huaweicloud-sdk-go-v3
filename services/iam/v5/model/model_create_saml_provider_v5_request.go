package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSamlProviderV5Request Request Object
type CreateSamlProviderV5Request struct {
	Body *CreateSamlProviderReqBody `json:"body,omitempty"`
}

func (o CreateSamlProviderV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSamlProviderV5Request struct{}"
	}

	return strings.Join([]string{"CreateSamlProviderV5Request", string(data)}, " ")
}
