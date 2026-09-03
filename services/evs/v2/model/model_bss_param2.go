package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BssParam2 包周期参数:  1. 包年包月的计量属性 2. 包年包月的数量
type BssParam2 struct {

	// 云硬盘的计费模式转成包周期后，是否自动支付
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o BssParam2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BssParam2 struct{}"
	}

	return strings.Join([]string{"BssParam2", string(data)}, " ")
}
