package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Validity struct {

	// 有效期类型，为必填值：   - **YEAR** : 年（12个月）   - **MONTH** : 月（统一按31天）   - **DAY** : 日   - **HOUR** : 小时
	Type string `json:"type"`

	// 证书有效期值，与type对应的类型值，换算成年需满足以下规则：   - 根CA，有效期小于等于30年；   - 从属CA与私有证书，有效期小于等于20年。
	Value int32 `json:"value"`

	// 起始时间，为可选值:   - 格式为时间戳（毫秒级），如1645146939688代表2022-02-18 09:15:39；   - 不早于当前时间5分钟，即start_from > (current_time - 5min)。
	StartFrom *int32 `json:"start_from,omitempty"`
}

func (o Validity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Validity struct{}"
	}

	return strings.Join([]string{"Validity", string(data)}, " ")
}
