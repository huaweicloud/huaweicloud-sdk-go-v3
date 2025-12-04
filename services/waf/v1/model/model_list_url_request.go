package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUrlRequest Request Object
type ListUrlRequest struct {

	// 受攻击次数最多的几条url
	Top int32 `json:"top"`

	// **参数解释：** 查询日志的时间范围，如1week（1周）、1month（1个月） **约束限制：** 不涉及 **取值范围：** - yesterday - today - 3days - 1week - 1month  **默认取值：** 不涉及
	Recent *string `json:"recent,omitempty"`

	// **参数解释：** 开始时间，统计周期的起始时间戳（毫秒级）。不使用recent参数时需要填写 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	From *int32 `json:"from,omitempty"`

	// **参数解释：** 结束时间，统计周期的终止时间戳（毫秒级）。不使用recent参数时需要填写 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	To *int32 `json:"to,omitempty"`

	// 要查询事件的域名id列表
	Hosts *[]string `json:"hosts,omitempty"`

	// 要查询事件的独享域名id列表
	Instances *[]string `json:"instances,omitempty"`
}

func (o ListUrlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUrlRequest struct{}"
	}

	return strings.Join([]string{"ListUrlRequest", string(data)}, " ")
}
