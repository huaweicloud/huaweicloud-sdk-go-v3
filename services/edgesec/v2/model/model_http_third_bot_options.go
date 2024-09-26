package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HttpThirdBotOptions 三方BOT操作
type HttpThirdBotOptions struct {
	RiverConfig *HttpRiverConfig `json:"river_config,omitempty"`
}

func (o HttpThirdBotOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpThirdBotOptions struct{}"
	}

	return strings.Join([]string{"HttpThirdBotOptions", string(data)}, " ")
}
