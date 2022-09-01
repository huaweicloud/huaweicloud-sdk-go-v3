package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FlavorAzObject struct {

	// 缓存容量（G Byte）。
	Capacity *string `json:"capacity,omitempty" xml:"capacity"`

	// 有资源的可用区编码。
	AzCodes *[]string `json:"az_codes,omitempty" xml:"az_codes"`
}

func (o FlavorAzObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorAzObject struct{}"
	}

	return strings.Join([]string{"FlavorAzObject", string(data)}, " ")
}
