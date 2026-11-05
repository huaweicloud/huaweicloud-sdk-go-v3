package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PayInfoBody 包周期集群相关参数，订购包周期集群时使用。
type PayInfoBody struct {

	// 订购周期类型。 - 2: 包月。 - 3: 包年。
	PayModel int32 `json:"payModel"`

	// 订购周期数。 - 若payModel为2，则有效值为1-9。 - 若payModel为3，则有效值为1-3。
	Period int32 `json:"period"`

	// 是否自动续订，为空时表示不自动续订。 - 1: 自动续订。 - 0：不自动续订（默认）。
	IsAutoRenew *int32 `json:"isAutoRenew,omitempty"`

	IsAutoPay *int32 `json:"isAutoPay,omitempty"`
}

func (o PayInfoBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PayInfoBody struct{}"
	}

	return strings.Join([]string{"PayInfoBody", string(data)}, " ")
}
