package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListVaultTopsResponseBodyVaults struct {

	// 存储库ID
	Id string `json:"id"`

	// 指标时间，单位毫秒
	Time int64 `json:"time"`

	// 指标数据维度
	Dimension string `json:"dimension"`

	// 指标数据取值
	Value float64 `json:"value"`

	// 指标数据的单位，例如：个、GB、%
	Unit *string `json:"unit,omitempty"`
}

func (o ListVaultTopsResponseBodyVaults) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVaultTopsResponseBodyVaults struct{}"
	}

	return strings.Join([]string{"ListVaultTopsResponseBodyVaults", string(data)}, " ")
}
