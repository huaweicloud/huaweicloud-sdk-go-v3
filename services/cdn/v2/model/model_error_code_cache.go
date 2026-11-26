package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ErrorCodeCache CDN状态码缓存时间。
type ErrorCodeCache struct {

	// **参数解释：** 需要缓存的错误码 **约束限制：** 不涉及 **取值范围：** - 3xx: 301, 302 - 4xx: 400, 401, 403, 404, 405, 407, 414, 416, 424, 451 - 5xx: 500, 501, 502, 503, 504, 509, 514  **默认取值：** 不涉及
	Code *int32 `json:"code,omitempty"`

	// **参数解释：** 错误码缓存时间 **约束限制：** 不涉及 **取值范围：** 0-31536000，单位：秒 > 3XX及416状态码缓存时间范围为0-20s  **默认取值：** 不涉及
	Ttl *int32 `json:"ttl,omitempty"`
}

func (o ErrorCodeCache) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ErrorCodeCache struct{}"
	}

	return strings.Join([]string{"ErrorCodeCache", string(data)}, " ")
}
