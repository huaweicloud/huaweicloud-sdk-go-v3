package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchResizeServersOption struct {

	// flavor
	FlavorRef string `json:"flavorRef"`

	Servers []ServerId `json:"servers"`

	CpuOptions *CpuOptions `json:"cpu_options,omitempty"`

	Mode *string `json:"mode,omitempty"`

	Promotion *Promotion `json:"promotion,omitempty"`
}

func (o BatchResizeServersOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResizeServersOption struct{}"
	}

	return strings.Join([]string{"BatchResizeServersOption", string(data)}, " ")
}
