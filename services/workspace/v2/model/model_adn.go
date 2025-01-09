package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Adn Adn信息。
type Adn struct {

	// adn带宽计费模式 - free：不计费。
	ChargeMode string `json:"charge_mode"`

	// 带宽大小，charge_mode为free时，不支持配置。
	BandwidthSize *int32 `json:"bandwidth_size,omitempty"`
}

func (o Adn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Adn struct{}"
	}

	return strings.Join([]string{"Adn", string(data)}, " ")
}
