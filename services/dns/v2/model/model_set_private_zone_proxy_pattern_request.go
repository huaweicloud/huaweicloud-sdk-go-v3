package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SetPrivateZoneProxyPatternRequest struct {

	// 待设置内网Zone的ID。
	ZoneId string `json:"zone_id"`

	Body *SetPrivateZoneProxyPatternRequestBody `json:"body,omitempty"`
}

func (o SetPrivateZoneProxyPatternRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetPrivateZoneProxyPatternRequest struct{}"
	}

	return strings.Join([]string{"SetPrivateZoneProxyPatternRequest", string(data)}, " ")
}
