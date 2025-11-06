package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetaData 查询结果元数据信息，包括分页信息等。
type MetaData struct {

	// **参数解释**： 当前返回结果条数。    **约束限制**： 不涉及。  **取值范围**： 在[0,2147483647]区间内 **默认取值**： 不涉及。
	Count int32 `json:"count"`

	// **参数解释**： 总条数。    **约束限制**： 不涉及。  **取值范围**： 在[0,2147483647]区间内 **默认取值**： 不涉及。
	Total int32 `json:"total"`

	// **参数解释**： 下一个开始的标记，用于分页。    **约束限制**： 不涉及。  **取值范围**： 长度为[1,9999]个数字 **默认取值**： 不涉及。
	Marker string `json:"marker"`
}

func (o MetaData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetaData struct{}"
	}

	return strings.Join([]string{"MetaData", string(data)}, " ")
}
