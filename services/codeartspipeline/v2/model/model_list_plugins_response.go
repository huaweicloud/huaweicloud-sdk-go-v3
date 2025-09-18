package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPluginsResponse Response Object
type ListPluginsResponse struct {

	// **参数解释**： 偏移量，表示从此偏移量开始查询。 **约束限制**： 不涉及。 **取值范围**： offset大于等于0。 **默认取值**： 默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 每次查询的条目数量。 **约束限制**： 不涉及。 **取值范围**： 大于等于0。 **默认取值**： 不涉及。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 总条目数量。 **取值范围**： 大于等于0。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 扩展插件列表。 **取值范围**： 不涉及。
	Data           *[]PluginBasicVo `json:"data,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListPluginsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPluginsResponse struct{}"
	}

	return strings.Join([]string{"ListPluginsResponse", string(data)}, " ")
}
