package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConcurrencyPackageUsingResponse Response Object
type ShowConcurrencyPackageUsingResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowConcurrencyPackageUsingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConcurrencyPackageUsingResponse struct{}"
	}

	return strings.Join([]string{"ShowConcurrencyPackageUsingResponse", string(data)}, " ")
}
