package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ErrorCodeCacheEngine **参数解释：** 将源站返回的错误状态码缓存到CDN节点，用户再次请求时由CDN直接响应给用户错误状态码 **约束限制：** 不涉及
type ErrorCodeCacheEngine struct {

	// **参数解释：** 需要缓存的错误码 **约束限制：** 不涉及 **取值范围：** - 3xx: 301, 302 - 4xx: 400, 403, 404, 405, 414 - 5xx: 501, 502, 503, 504 **默认取值：** 不涉及
	Code int32 `json:"code"`

	// **参数解释：** 错误码缓存时间 **约束限制：** 不涉及 **取值范围：** 0-31536000，单位：秒 > 3XX状态码缓存时间范围为0-20s  **默认取值：** 不涉及
	Ttl int32 `json:"ttl"`
}

func (o ErrorCodeCacheEngine) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorCodeCacheEngine struct{}"
	}

	return strings.Join([]string{"ErrorCodeCacheEngine", string(data)}, " ")
}
