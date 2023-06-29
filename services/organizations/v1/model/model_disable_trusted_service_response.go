package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableTrustedServiceResponse Response Object
type DisableTrustedServiceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisableTrustedServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableTrustedServiceResponse struct{}"
	}

	return strings.Join([]string{"DisableTrustedServiceResponse", string(data)}, " ")
}
