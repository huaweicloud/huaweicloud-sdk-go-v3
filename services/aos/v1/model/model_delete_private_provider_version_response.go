package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePrivateProviderVersionResponse Response Object
type DeletePrivateProviderVersionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePrivateProviderVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateProviderVersionResponse struct{}"
	}

	return strings.Join([]string{"DeletePrivateProviderVersionResponse", string(data)}, " ")
}
