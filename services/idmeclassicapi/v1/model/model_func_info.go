package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FuncInfo struct {

	// **参数解释：**  指定简单函数名称。  **约束限制：**  不涉及。  **取值范围：**  - AVG：求平均值。 - COUNT：求总数。 - MAX：求最大值。 - MIX：求最小值。  **默认取值：**  不涉及。
	Func string `json:"func"`

	// **参数解释：**  指定简单函数以哪个属性为维度操作。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	FuncBy string `json:"funcBy"`
}

func (o FuncInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FuncInfo struct{}"
	}

	return strings.Join([]string{"FuncInfo", string(data)}, " ")
}
