package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComputingSpecDto struct {

	// 规格编号
	Code string `json:"code"`

	// 规格名称
	Name string `json:"name"`

	// 内存
	Ram *int64 `json:"ram,omitempty"`

	// vcpus
	Vcpus *string `json:"vcpus,omitempty"`

	// 最大带宽
	MaxRate *string `json:"max_rate,omitempty"`

	// 基准带宽
	MinRate *string `json:"min_rate,omitempty"`

	// 最大收发包能力
	MaxPps *string `json:"max_pps,omitempty"`
}

func (o ComputingSpecDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComputingSpecDto struct{}"
	}

	return strings.Join([]string{"ComputingSpecDto", string(data)}, " ")
}
