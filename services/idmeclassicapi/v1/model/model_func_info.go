package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FuncInfo struct {

	// **参数解释：**  指定聚合函数名称，用于定义对数据执行的统计计算方式。  **约束限制：**  不涉及。  **取值范围：**  - AVG：求平均值。 - COUNT：求总数。 - MAX：求最大值。 - MIN：求最小值。  **默认取值：**  不涉及。
	Func string `json:"func"`

	// **参数解释：**  指定聚合函数的操作维度属性名称，即对哪个数据属性进行聚合计算。 例如func为AVG且funcBy为“currentTemperature”时，表示计算平均运行温度。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	FuncBy string `json:"funcBy"`
}

func (o FuncInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FuncInfo struct{}"
	}

	return strings.Join([]string{"FuncInfo", string(data)}, " ")
}
