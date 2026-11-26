package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BrowserCacheRulesCondition **参数解释：** 匹配条件 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type BrowserCacheRulesCondition struct {

	// **参数解释：** 匹配类型 **约束限制：** 不涉及 **取值范围：** - all: 匹配所有文件 - file_extension: 按文件后缀匹配 - catalog: 按目录匹配 - full_path: 全路径匹配 - home_page: 按首页匹配  **默认取值：** 不涉及
	MatchType string `json:"match_type"`

	// **参数解释：** 缓存匹配设置 **约束限制：** - 当match_type为all时，置空 - 当match_type为file_extension时，为文件后缀，输入首字符为“.”，以“,”进行分隔， 如.jpg,.zip,.exe，并且输入的文 件名后缀总数不超过20个 - 当match_type为catalog时，为目录，输入要求以“/”作为首字符， 以“,”进行分隔，如/test/folder01,/test/folder02，并且输入的目录路径总数不超过20个 - 当match_type为full_path时，为全路径，输入要求以“/”作为首字符，支持匹配指定目录下的具体文件，或者带通配符“\\*”的文件， “\\*”的位置必须在最后一个“/”的后一位，且不能以“*”结尾，单条全路径缓存规则里仅支持配置一个全路径，如/test/index.html或/test/\\*.jpg - 当match_type为home_page时，置空  **取值范围：** 不涉及 **默认取值：** 不涉及
	MatchValue *string `json:"match_value,omitempty"`

	// **参数解释：** 浏览器缓存的优先级, 数值越大优先级越高 **约束限制：** 不支持多条规则设置同一优先级 **取值范围：** - 1-1000  **默认取值：** 不涉及
	Priority int32 `json:"priority"`
}

func (o BrowserCacheRulesCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BrowserCacheRulesCondition struct{}"
	}

	return strings.Join([]string{"BrowserCacheRulesCondition", string(data)}, " ")
}
