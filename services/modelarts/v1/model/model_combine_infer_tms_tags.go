package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CombineInferTmsTags 多标签相同key合并value的数据结构。
type CombineInferTmsTags struct {

	// **参数解释：** 标签的key。 **取值范围：** 不涉及。
	Key string `json:"key"`

	// **参数解释：** 相同key的标签value合并后的列表。 **取值范围：** 不涉及。
	Values []string `json:"values"`
}

func (o CombineInferTmsTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CombineInferTmsTags struct{}"
	}

	return strings.Join([]string{"CombineInferTmsTags", string(data)}, " ")
}
