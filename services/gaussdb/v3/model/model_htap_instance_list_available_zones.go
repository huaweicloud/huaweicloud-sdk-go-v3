package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HtapInstanceListAvailableZones struct {

	// 可用区码。
	Code *string `json:"code,omitempty"`

	// 可用区描述。
	Description *string `json:"description,omitempty"`

	// 可用区类型。
	AzType *string `json:"az_type,omitempty"`
}

func (o HtapInstanceListAvailableZones) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HtapInstanceListAvailableZones struct{}"
	}

	return strings.Join([]string{"HtapInstanceListAvailableZones", string(data)}, " ")
}
