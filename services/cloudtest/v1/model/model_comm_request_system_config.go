package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommRequestSystemConfig struct {
	Params *SystemConfig `json:"params"`
}

func (o CommRequestSystemConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommRequestSystemConfig struct{}"
	}

	return strings.Join([]string{"CommRequestSystemConfig", string(data)}, " ")
}
