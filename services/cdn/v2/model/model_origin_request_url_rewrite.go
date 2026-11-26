package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OriginRequestUrlRewrite **参数解释：** 改写回源URL **约束限制：** 最多配置20条
type OriginRequestUrlRewrite struct {

	// **参数解释：** 回源URL改写规则的优先级 > 数值越大优先级越高，CDN按照优先级高的规则执行URL改写  **约束限制：** - 不支持多条回源URL改写规则设置同一优先级 - 优先级不能输入为空  **取值范围：** 1-1000 **默认取值：** 不涉及
	Priority int32 `json:"priority"`

	// **参数解释：** 匹配类型 **约束限制：** 不涉及 **取值范围：** - all: 所有文件 - file_path: URI路径 - wildcard: 通配符 - full_path: 全路径  **默认取值：** 不涉及
	MatchType string `json:"match_type"`

	// **参数解释：** 需要替换的URI **约束限制：** 匹配方式为“所有文件”时，不支持配置参数 **取值范围：** - 1-512个字符 - 支持通配符\\*匹配，如：/test/\\*_/\\*.mp4 - 以正斜线（/）开头的URI，不含http(s)://头及域名  **默认取值：** 不涉及
	SourceUrl *string `json:"source_url,omitempty"`

	// **参数解释：** 替换后的URI **约束限制：** **取值范围：** - 1-256个字符 - 以正斜线（/）开头的URI，不含http(s)://头及域名  > 通配符 * 可通过$n捕获（n=1,2,3...，例如：/newtest/$1/$2.jpg）  **默认取值：** 不涉及
	TargetUrl string `json:"target_url"`
}

func (o OriginRequestUrlRewrite) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OriginRequestUrlRewrite struct{}"
	}

	return strings.Join([]string{"OriginRequestUrlRewrite", string(data)}, " ")
}
