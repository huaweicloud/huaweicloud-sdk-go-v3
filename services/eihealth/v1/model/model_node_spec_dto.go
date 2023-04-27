package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeSpecDto struct {

	// 规格编号
	Code *string `json:"code,omitempty"`

	// 规格名称
	Name *string `json:"name,omitempty"`

	// 内存
	Ram *int64 `json:"ram,omitempty"`

	// vcpus
	Vcpus *string `json:"vcpus,omitempty"`
}

func (o NodeSpecDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeSpecDto struct{}"
	}

	return strings.Join([]string{"NodeSpecDto", string(data)}, " ")
}
