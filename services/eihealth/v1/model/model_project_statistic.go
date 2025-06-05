package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProjectStatistic 空间资源统计信息
type ProjectStatistic struct {

	// 空间ID。
	Id *string `json:"id,omitempty"`

	// 空间名称。
	Name *string `json:"name,omitempty"`

	// 用户所属空间的角色。
	RoleType *string `json:"role_type,omitempty"`

	// 空间所有者。
	Creator *string `json:"creator,omitempty"`

	// 空间创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 空间更新时间。
	UpdateTime *string `json:"update_time,omitempty"`

	// 空间置顶时间。
	TopTime *string `json:"top_time,omitempty"`

	// 用户资源统计详情
	UserStatistics *[]StatisticDto `json:"user_statistics,omitempty"`

	// 总资源统计详情
	TotalStatistics *[]StatisticDto `json:"total_statistics,omitempty"`
}

func (o ProjectStatistic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectStatistic struct{}"
	}

	return strings.Join([]string{"ProjectStatistic", string(data)}, " ")
}
