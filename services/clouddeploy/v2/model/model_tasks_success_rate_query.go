package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询多个应用的部署成功率的请求体
type TasksSuccessRateQuery struct {

	// 部署应用开始时间范围的左边界（包含），格式yyyy-MM-dd
	StartDate string `json:"start_date"`

	// 部署应用开始时间范围的右边界（包含），格式yyyy-MM-dd 。最大时间范围为1年。
	EndDate string `json:"end_date"`

	// 任务id列表
	TaskIds []string `json:"task_ids"`
}

func (o TasksSuccessRateQuery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TasksSuccessRateQuery struct{}"
	}

	return strings.Join([]string{"TasksSuccessRateQuery", string(data)}, " ")
}
