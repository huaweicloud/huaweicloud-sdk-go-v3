package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SelectedUnit **参数解释**： 用户在页面中选择的指标单位， 用于后续指标数据回显和计算。字符串最大长度为64。    **约束限制**： 不涉及。 **取值范围**： 长度为[0,64]个字符。        **默认取值**： 不涉及。
type SelectedUnit struct {
}

func (o SelectedUnit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SelectedUnit struct{}"
	}

	return strings.Join([]string{"SelectedUnit", string(data)}, " ")
}
