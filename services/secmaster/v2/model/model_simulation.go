package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Simulation 是否仿真
type Simulation struct {
}

func (o Simulation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Simulation struct{}"
	}

	return strings.Join([]string{"Simulation", string(data)}, " ")
}
