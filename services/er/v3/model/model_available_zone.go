package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AvailableZone struct {

	// 可用区编码
	Code *string `json:"code,omitempty"`

	// 可用区是否可用available|unavailable
	State *string `json:"state,omitempty"`
}

func (o AvailableZone) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AvailableZone struct{}"
	}

	return strings.Join([]string{"AvailableZone", string(data)}, " ")
}
