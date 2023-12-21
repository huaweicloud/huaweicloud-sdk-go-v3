package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FuncInfo struct {

	// 指定简单函数名称。 - AVG：求平均值。 - COUNT：求总数。 - MAX：求最大值。 - MIX：求最小值。
	Func string `json:"func"`

	// 指定简单函数以哪个属性为维度操作。
	FuncBy string `json:"funcBy"`
}

func (o FuncInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FuncInfo struct{}"
	}

	return strings.Join([]string{"FuncInfo", string(data)}, " ")
}
