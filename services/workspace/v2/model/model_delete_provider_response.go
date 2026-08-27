package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProviderResponse Response Object
type DeleteProviderResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteProviderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProviderResponse struct{}"
	}

	return strings.Join([]string{"DeleteProviderResponse", string(data)}, " ")
}
