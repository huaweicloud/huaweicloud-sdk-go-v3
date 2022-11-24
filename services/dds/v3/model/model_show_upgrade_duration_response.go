package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowUpgradeDurationResponse struct {

	// 升级策略列表
	Strategies     *[]DurationStrategies `json:"strategies,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowUpgradeDurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUpgradeDurationResponse struct{}"
	}

	return strings.Join([]string{"ShowUpgradeDurationResponse", string(data)}, " ")
}
