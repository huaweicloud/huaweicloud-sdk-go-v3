package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccessLogging struct {

	// LTS配置
	Lts *[]LtsConfig `json:"lts,omitempty"`
}

func (o AccessLogging) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessLogging struct{}"
	}

	return strings.Join([]string{"AccessLogging", string(data)}, " ")
}
