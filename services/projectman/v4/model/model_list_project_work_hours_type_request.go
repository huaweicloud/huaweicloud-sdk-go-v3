package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListProjectWorkHoursTypeRequest struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 每页显示的数量，默认显示10条，最多显示100条
	Limit *int32 `json:"limit,omitempty"`

	// 分页索引，偏移量offset是limit的整数倍，limit=10,offset=0,10,20...
	Offset *int32 `json:"offset,omitempty"`

	// 工时类型状态，支持按状态筛选查询，置空时查询所有工时类型，1表示查询启用的工时类型，2表示查询未启用的工时类型
	Status *int32 `json:"status,omitempty"`
}

func (o ListProjectWorkHoursTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectWorkHoursTypeRequest struct{}"
	}

	return strings.Join([]string{"ListProjectWorkHoursTypeRequest", string(data)}, " ")
}
