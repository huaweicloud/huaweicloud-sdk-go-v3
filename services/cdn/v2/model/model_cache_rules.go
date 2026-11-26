package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CacheRules 缓存规则
type CacheRules struct {

	// **参数解释：** 匹配类型 **约束限制：** 不涉及 **取值范围：** - all: 匹配所有文件 - file_extension: 按文件后缀匹配 - catalog: 按目录匹配 - full_path: 全路径匹配 - home_page: 按首页匹配  **默认取值：** 不涉及
	MatchType string `json:"match_type"`

	// **参数解释：** 缓存规则匹配内容 **约束限制：** - 当match_type为all时，置空 - 当match_type为file_extension时，为文件后缀，输入首字符为“.”，以“,”或“;”分隔，如.jpg,.zip,.exe，并且输入的文件名后缀总数不超过100个，总字符不超过1000 - 当match_type为catalog时，为目录路径，输入要求以“/”作为首字符， 以“,”或“;”分隔，如/test/folder01,/test/folder02，并且输入的目录路径总数不超过20个，总字符不超过255 - 当match_type为full_path时，为全路径，输入要求以“/”作为首字符，支持匹配指定目录下的具体文件，或者带通配符“\\*”的文件，单条全路径缓存规则里仅支持配置一个全路径，如/test/index.html或/test/\\*.jpg，不可配置/_* - 当match_type为home_page时，置空  **取值范围：** 不涉及 **默认取值：** 不涉及
	MatchValue *string `json:"match_value,omitempty"`

	// **参数解释：** 资源在CDN节点的缓存过期时间，单位由ttl_unit参数指定 **约束限制：** 最大支持365天 **取值范围：** 不涉及 **默认取值：** 0
	Ttl *int32 `json:"ttl,omitempty"`

	// **参数解释：** 缓存过期时间单位 **约束限制：** 不涉及 **取值范围：** - s: 秒 - m: 分 - h: 小时 - d: 天  **默认取值：** 不涉及
	TtlUnit string `json:"ttl_unit"`

	// **参数解释：** 缓存规则的优先级，数值越大，优先级越高 **约束限制：** 优先级不能相同 **取值范围：** 1-1000 **默认取值：** 不涉及
	Priority int32 `json:"priority"`

	// **参数解释：** 缓存过期时间来源，设置CDN节点的缓存遵循源站还是CDN侧的配置 **约束限制：** 不涉及 **取值范围：** - on: CDN节点的缓存过期时间遵循源站的设置 - off: CDN节点的缓存过期时间遵循“缓存规则”中的“缓存过期时间” - min_ttl: CDN节点的缓存过期时间取缓存规则和源站二者的最小值  **默认取值：** off: CDN节点的缓存过期时间遵循“缓存规则”中的“缓存过期时间”
	FollowOrigin *string `json:"follow_origin,omitempty"`

	// **参数解释：** 强制缓存：CDN节点缓存过期时间是否忽略源站响应头Cache-Control中的no-cache、private、no-store字段 **约束限制：** 强制缓存与缓存过期时间来源功能配合使用，具体使用限制及配置效果请参考CDN用户指南的配置节点缓存规则章节 **取值范围：** - on: 打开强制缓存 - off: 关闭强制缓存  **默认取值：** off: 关闭强制缓存
	ForceCache *string `json:"force_cache,omitempty"`

	// **参数解释：** URL参数 **约束限制：** 不涉及 **取值范围：** - del_params: 忽略指定URL参数 - reserve_params: 保留指定URL参数 - ignore_url_params: 忽略全部URL参数 - full_url: 使用完整URL参数  **默认取值：** full_url: 使用完整URL参数
	UrlParameterType *string `json:"url_parameter_type,omitempty"`

	// **参数解释：** URL参数值 **约束限制：** - 最多配置30条参数值 - 以“,”或“;”分隔 - 当url_parameter_type为del_params或reserve_params时必填  **取值范围：** 不涉及 **默认取值：** 不涉及
	UrlParameterValue *string `json:"url_parameter_value,omitempty"`

	// **参数解释：** 如果您的源站的Cache-Control头部设置了stale-while-revalidate=***（***为时间），可以在CDN侧开启SWR开关，当CDN节点缓存的资源过期后，如果客户端请求该资源时过期时间没有超过stale-while-revalidate设置的时间，浏览器仍然返回已缓存的资源，同时cdn节点会回源请求最新的资源并缓存，用户再次请求时就会得到最新的资源。 **约束限制：** 不涉及 **取值范围：** - on: 开启 - off: 关闭  **默认取值：** off: 关闭
	StaleWhileRevalidate *string `json:"stale_while_revalidate,omitempty"`
}

func (o CacheRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CacheRules struct{}"
	}

	return strings.Join([]string{"CacheRules", string(data)}, " ")
}
