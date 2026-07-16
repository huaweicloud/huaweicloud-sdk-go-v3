package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Npu Ascend规格信息。
type Npu struct {

	// npu卡数。
	UnitNum *string `json:"unit_num,omitempty"`

	// 产品名。
	ProductName *string `json:"product_name,omitempty"`

	// 内存。
	Memory *string `json:"memory,omitempty"`
}

func (o Npu) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Npu struct{}"
	}

	return strings.Join([]string{"Npu", string(data)}, " ")
}
