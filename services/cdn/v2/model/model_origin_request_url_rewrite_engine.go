package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OriginRequestUrlRewriteEngine **参数解释：** 改写回源URL **约束限制：** 最多配置20条
type OriginRequestUrlRewriteEngine struct {

	// **参数解释：** 改写方式 **约束限制：** 不涉及 **取值范围：** - simple: 精确改写 - wildcard: 捕获改写 - regex: 正则改写（白名单功能，请提交工单开放该配置） **默认取值：** 不涉及
	RewriteType string `json:"rewrite_type"`

	// **参数解释：** 需要替换的URI **约束限制：** 当rewrite_type为wildcard或regex时，该参数必填 当rewrite_type为regex时，该参数必填必须以“^/”开始，如：^/test **取值范围：** - 1-512个字符 - 支持通配符\\*匹配，如：/test/\\*_/\\*.mp4 - 以正斜线（/）开头的URI，不含http(s)://头及域名 **默认取值：** 不涉及
	SourceUrl *string `json:"source_url,omitempty"`

	// **参数解释：** 替换后的URI **约束限制：** **取值范围：** - 1-256个字符 - 以正斜线（/）开头的URI，不含http(s)://头及域名  > 通配符 * 可通过$n捕获（n=1,2,3...，例如：/newtest/$1/$2.jpg）  **默认取值：** 不涉及
	TargetUrl string `json:"target_url"`
}

func (o OriginRequestUrlRewriteEngine) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OriginRequestUrlRewriteEngine struct{}"
	}

	return strings.Join([]string{"OriginRequestUrlRewriteEngine", string(data)}, " ")
}
