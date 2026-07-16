package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Cpu cpu规格信息。
type Cpu struct {

	// cpu架构。
	Arch *string `json:"arch,omitempty"`

	// 核数。
	CoreNum *int32 `json:"core_num,omitempty"`
}

func (o Cpu) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Cpu struct{}"
	}

	return strings.Join([]string{"Cpu", string(data)}, " ")
}
