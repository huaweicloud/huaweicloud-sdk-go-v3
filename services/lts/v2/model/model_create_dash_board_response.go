package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDashBoardResponse Response Object
type CreateDashBoardResponse struct {

	// 仪表盘图表
	Charts *[]string `json:"charts,omitempty"`

	// 过滤条件
	Filters *[]string `json:"filters,omitempty"`

	// 日志组名称
	GroupName *string `json:"group_name,omitempty"`

	// 仪表盘id
	Id *string `json:"id,omitempty"`

	// 最近修改时间
	LastUpdateTime *int64 `json:"last_update_time,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 仪表盘名称
	Title *string `json:"title,omitempty"`

	// 是否使用模板
	UseSystemTemplate *bool `json:"useSystemTemplate,omitempty"`

	// 纯净模式开关
	ConciseModeEnable *bool `json:"concise_mode_enable,omitempty"`

	// 仪表盘标签
	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateDashBoardResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDashBoardResponse struct{}"
	}

	return strings.Join([]string{"CreateDashBoardResponse", string(data)}, " ")
}
