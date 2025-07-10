package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CacheRulesEngine **参数解释：** 通过不同参数控制源站资源在CDN节点的缓存时长 **约束限制：** 不涉及
type CacheRulesEngine struct {

	// **参数解释：** 资源在CDN节点的缓存过期时间 **约束限制：** 最大支持365天 **取值范围：** 不涉及 **默认取值：** 不涉及
	Ttl int32 `json:"ttl"`

	// **参数解释：** 缓存过期时间单位 **约束限制：** 不涉及 **取值范围：** - s: 秒 - m: 分 - h: 小时 - d: 天 **默认取值：** 不涉及
	TtlUnit string `json:"ttl_unit"`

	// **参数解释：** 缓存过期时间来源，设置CDN节点的缓存遵循源站还是CDN侧的配置 **约束限制：** 不涉及 **取值范围：** - on: CDN节点的缓存过期时间遵循源站的设置 - off: CDN节点的缓存过期时间遵循“缓存规则”中的“缓存过期时间” - min_ttl: CDN节点的缓存过期时间取缓存规则和源站二者的最小值 **默认取值：** off: CDN节点的缓存过期时间遵循“缓存规则”中的“缓存过期时间”
	FollowOrigin string `json:"follow_origin"`

	// **参数解释：** 强制缓存：CDN节点缓存过期时间是否忽略源站响应头Cache-Control中的no-cache、private、no-store字段 **约束限制：** 强制缓存与缓存过期时间来源功能配合使用，具体使用限制及配置效果请参考CDN用户指南的配置节点缓存规则章节 **取值范围：** - on: 打开强制缓存 - off: 关闭强制缓存 **默认取值：** off: 关闭强制缓存
	ForceCache *string `json:"force_cache,omitempty"`
}

func (o CacheRulesEngine) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CacheRulesEngine struct{}"
	}

	return strings.Join([]string{"CacheRulesEngine", string(data)}, " ")
}
