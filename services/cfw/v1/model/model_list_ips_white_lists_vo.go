package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIpsWhiteListsVo **参数解释**：  IPS白名单响应体 **取值范围**： 不涉及
type ListIpsWhiteListsVo struct {

	// **参数解释**：  每页数量 **取值范围**： 不涉及
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：  偏移量 **取值范围**： 不涉及
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：  总数 **取值范围**： 不涉及
	Total *int64 `json:"total,omitempty"`

	// **参数解释**：  IPS白名单列表 **取值范围**： 不涉及
	Records *[]IpsWhiteListVo `json:"records,omitempty"`
}

func (o ListIpsWhiteListsVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpsWhiteListsVo struct{}"
	}

	return strings.Join([]string{"ListIpsWhiteListsVo", string(data)}, " ")
}
