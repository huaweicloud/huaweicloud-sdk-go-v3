package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListThreatsRequest Request Object
type ListThreatsRequest struct {

	// **参数解释：** 查询日志的时间范围，如1week（1周）、1month（1个月） **约束限制：** 不涉及 **取值范围：** - yesterday - today - 3days - 1week - 1month  **默认取值：** 不涉及
	Recent string `json:"recent"`

	// 要查询事件的域名列表
	Hosts *[]string `json:"hosts,omitempty"`
}

func (o ListThreatsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListThreatsRequest struct{}"
	}

	return strings.Join([]string{"ListThreatsRequest", string(data)}, " ")
}
