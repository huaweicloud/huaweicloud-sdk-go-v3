package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFactoryTaskOverviewRequest Request Object
type ListFactoryTaskOverviewRequest struct {

	// 工作空间ID
	Workspace *string `json:"workspace,omitempty"`

	// 是否查询当前用户的实例，默认为false，表示查询全部用户实例，为true时，表示查询当前用户的实例。
	IsOwn *string `json:"is_own,omitempty"`

	// 查询的天数，取值范围为：today、yesterday、before_yesterday、all，默认为today，表示查询今天的数据，支持查询近7天的数据。 today：查询当天的实例状态数量， yesterday：查询昨天的实例状态数量， before_yesterday：查询前天的实例状态数量， all：查询7天前到当天的实例状态总量。
	QueryDays *string `json:"query_days,omitempty"`
}

func (o ListFactoryTaskOverviewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactoryTaskOverviewRequest struct{}"
	}

	return strings.Join([]string{"ListFactoryTaskOverviewRequest", string(data)}, " ")
}
