package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlexibleOrigins **参数解释：** 高级回源，实现根据不同的资源类型或路径回源到不同源站 **约束限制：** 最多配置20条
type FlexibleOrigins struct {

	// **参数解释：** URI的匹配方式 **约束限制：** 不涉及 **取值范围：** - file_extension: 文件后缀匹配 - file_path: 目录路径 - all: 匹配所有文件  **默认取值：** 不涉及
	MatchType string `json:"match_type"`

	// **参数解释：** URI的匹配规则 **约束限制：** match_type为file_extension时：   - 支持所有格式的文件类型   - 输入首字符为“.”，以“;”进行分隔   - 输入的文件后缀名总数不能超过20个  match_type为file_path时：   - 要求以“/”作为首字符   - 以“;”进行分隔   - 输入的目录路径总数不能超过20个  match_type为all时：置空  > 文件名后缀英文字符大、小写代表不同资源  **取值范围：** 1-512个字符 **默认取值：** 不涉及
	MatchPattern *string `json:"match_pattern,omitempty"`

	// **参数解释：** 优先级 > 数值越大优先级越高  **约束限制：** 不涉及 **取值范围：** 1-1000 **默认取值：** 不涉及
	Priority int32 `json:"priority"`

	// **参数解释：** 配置客户端请求匹配当前规则时，CDN节点回源访问的源站 **约束限制：** 每条规则的回源源站数量不超过1个 **取值范围：** 不涉及 **默认取值：** 不涉及
	BackSources []BackSources `json:"back_sources"`
}

func (o FlexibleOrigins) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlexibleOrigins struct{}"
	}

	return strings.Join([]string{"FlexibleOrigins", string(data)}, " ")
}
