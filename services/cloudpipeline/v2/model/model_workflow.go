package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 流水线参数详情
type Workflow struct {

	// 任务类型,list类型数据
	Parameter []PipelineParam `json:"parameter" xml:"parameter"`

	// 源码仓,list类型数据
	Source []Source `json:"source" xml:"source"`

	// 流水线名字
	Name string `json:"name" xml:"name"`

	// 项目ID
	ProjectId string `json:"project_id" xml:"project_id"`

	// 项目名字
	ProjectName string `json:"project_name" xml:"project_name"`
}

func (o Workflow) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Workflow struct{}"
	}

	return strings.Join([]string{"Workflow", string(data)}, " ")
}
