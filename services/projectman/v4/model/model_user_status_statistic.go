package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UserStatusStatistic struct {
	User *IssueUser `json:"user,omitempty"`

	// 满足条件的工作项总数
	ItemCount *int32 `json:"item_count,omitempty"`

	// 工作项对应状态的统计计数
	Data map[string]int32 `json:"data,omitempty"`
}

func (o UserStatusStatistic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserStatusStatistic struct{}"
	}

	return strings.Join([]string{"UserStatusStatistic", string(data)}, " ")
}
