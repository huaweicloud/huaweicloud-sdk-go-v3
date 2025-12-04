package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSourceIpTop5Request Request Object
type ListSourceIpTop5Request struct {

	// **参数解释：** 查询日志的时间范围，如1week（1周）、1month（1个月） **约束限制：** 不涉及 **取值范围：** - yesterday - today - 3days - 1week - 1month  **默认取值：** 不涉及
	Recent string `json:"recent"`

	// 要查询事件域名id列表
	Hosts []string `json:"hosts"`
}

func (o ListSourceIpTop5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSourceIpTop5Request struct{}"
	}

	return strings.Join([]string{"ListSourceIpTop5Request", string(data)}, " ")
}
