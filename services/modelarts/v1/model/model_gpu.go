package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Gpu gpu规格信息。
type Gpu struct {

	// gpu卡数。
	UnitNum *int32 `json:"unit_num,omitempty"`

	// 产品名。
	ProductName *string `json:"product_name,omitempty"`

	// 内存。
	Memory *string `json:"memory,omitempty"`
}

func (o Gpu) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Gpu struct{}"
	}

	return strings.Join([]string{"Gpu", string(data)}, " ")
}
