package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FlavorAzObject struct {
	// 缓存容量（G Byte）。

	Capacity *string `json:"capacity,omitempty"`
	// 有资源的可用区编码。

	AzCodes *[]string `json:"az_codes,omitempty"`
}

func (o FlavorAzObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorAzObject struct{}"
	}

	return strings.Join([]string{"FlavorAzObject", string(data)}, " ")
}
