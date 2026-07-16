package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RateLimitResponse **参数解释：** 流量控制配置。
type RateLimitResponse struct {

	// **参数解释：** 服务流量限制是指指定单位内一个服务能够被访问的次数上限。 **取值范围：** 1-10000。
	Num int32 `json:"num"`

	// **参数解释：** 流量限制单元。 **取值范围：** - NANOS（纳秒）。 - MICROS（微秒）。 - MILLIS（毫秒）。 - SECONDS（秒）。 - MINUTES（分钟）。 - HOURS（小时）。 - DAYS（天）。
	Unit string `json:"unit"`
}

func (o RateLimitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RateLimitResponse struct{}"
	}

	return strings.Join([]string{"RateLimitResponse", string(data)}, " ")
}
