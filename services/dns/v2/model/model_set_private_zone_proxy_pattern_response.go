package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SetPrivateZoneProxyPatternResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetPrivateZoneProxyPatternResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetPrivateZoneProxyPatternResponse struct{}"
	}

	return strings.Join([]string{"SetPrivateZoneProxyPatternResponse", string(data)}, " ")
}
