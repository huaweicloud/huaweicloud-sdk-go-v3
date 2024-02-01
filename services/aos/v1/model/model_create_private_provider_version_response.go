package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePrivateProviderVersionResponse Response Object
type CreatePrivateProviderVersionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreatePrivateProviderVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateProviderVersionResponse struct{}"
	}

	return strings.Join([]string{"CreatePrivateProviderVersionResponse", string(data)}, " ")
}
