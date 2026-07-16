package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServiceLimit **参数解释：** 服务限制信息。 **约束限制：** 不涉及。
type ServiceLimit struct {
	RateLimit *RateLimit `json:"rate_limit"`

	// **参数解释：** 请求大小限制。 **约束限制：** 不涉及。 **取值范围：** 1-50M。 **默认取值：** 不涉及。
	RequestSizeLimit int32 `json:"request_size_limit"`

	// **参数解释：** 超时时间。 **约束限制：** 不涉及。 **取值范围：** 1到7200秒。 **默认取值：** 不涉及。
	RequestTimeout int32 `json:"request_timeout"`

	// **参数解释：** IP白名单。 **约束限制：** 不涉及。
	IpWhiteList *[]string `json:"ip_white_list,omitempty"`

	// **参数解释：** IP黑名单。 **约束限制：** 不涉及。
	IpBlackList *[]string `json:"ip_black_list,omitempty"`
}

func (o ServiceLimit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceLimit struct{}"
	}

	return strings.Join([]string{"ServiceLimit", string(data)}, " ")
}
