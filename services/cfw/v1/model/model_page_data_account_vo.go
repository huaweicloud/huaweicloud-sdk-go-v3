package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageDataAccountVo struct {

	// **参数解释**： 数量限制 **取值范围**： 不涉及
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 偏移量 **取值范围**： 不涉及
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 账号列表 **取值范围**： 不涉及
	Records *[]AccountVo `json:"records,omitempty"`

	// **参数解释**： 总数 **取值范围**： 不涉及
	Total *int64 `json:"total,omitempty"`
}

func (o PageDataAccountVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageDataAccountVo struct{}"
	}

	return strings.Join([]string{"PageDataAccountVo", string(data)}, " ")
}
