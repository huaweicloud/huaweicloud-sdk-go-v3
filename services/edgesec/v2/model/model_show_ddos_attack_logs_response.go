package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDdosAttackLogsResponse Response Object
type ShowDdosAttackLogsResponse struct {

	// 攻击日志数量
	Total *int64 `json:"total,omitempty"`

	// 攻击日志详情
	Items          *[]DdosAttackLog `json:"items,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowDdosAttackLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDdosAttackLogsResponse struct{}"
	}

	return strings.Join([]string{"ShowDdosAttackLogsResponse", string(data)}, " ")
}
