package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServiceLimitResponse **参数解释：** 服务限制信息。
type ServiceLimitResponse struct {
	RateLimit *RateLimitResponse `json:"rate_limit"`

	// **参数解释：** 请求大小限制。 **取值范围：** 1-50M。
	RequestSizeLimit int32 `json:"request_size_limit"`

	// **参数解释：** 超时时间。 **取值范围：** 1到7200秒。
	RequestTimeout int32 `json:"request_timeout"`

	// **参数解释：** IP白名单。
	IpWhiteList *[]string `json:"ip_white_list,omitempty"`

	// **参数解释：** IP黑名单。
	IpBlackList *[]string `json:"ip_black_list,omitempty"`
}

func (o ServiceLimitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceLimitResponse struct{}"
	}

	return strings.Join([]string{"ServiceLimitResponse", string(data)}, " ")
}
