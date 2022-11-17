package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListEastWestFirewallRequest struct {

	// 每页显示个数
	Limit int32 `json:"limit"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset int32 `json:"offset"`
}

func (o ListEastWestFirewallRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEastWestFirewallRequest struct{}"
	}

	return strings.Join([]string{"ListEastWestFirewallRequest", string(data)}, " ")
}
