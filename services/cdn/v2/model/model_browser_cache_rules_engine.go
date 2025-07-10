package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BrowserCacheRulesEngine **参数解释：** 浏览器缓存过期时间，当终端用户请求资源时，如果浏览器有缓存，直接返回给用户 **约束限制：** 不涉及
type BrowserCacheRulesEngine struct {

	// **参数解释：** 缓存生效类型 **约束限制：** 不涉及 **取值范围：** - follow_origin: 遵循源站的缓存策略，即Cache-Control头部的设置 - ttl: 浏览器缓存遵循当前规则设置的过期时间 - never: 浏览器不缓存资源 **默认取值：** 不涉及
	CacheType string `json:"cache_type"`

	// **参数解释：** 缓存过期时间 **约束限制：** - 最大支持365天 - 当缓存生效类型为ttl时必填 **取值范围：** 不涉及 **默认取值：** 不涉及
	Ttl *int32 `json:"ttl,omitempty"`

	// **参数解释：** 缓存过期时间单位 **约束限制：** 当缓存生效类型为ttl时必填 **取值范围：** - s：秒 - m：分种 - h：小时 - d：天 **默认取值：** 不涉及
	TtlUnit *string `json:"ttl_unit,omitempty"`
}

func (o BrowserCacheRulesEngine) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BrowserCacheRulesEngine struct{}"
	}

	return strings.Join([]string{"BrowserCacheRulesEngine", string(data)}, " ")
}
