/*
 * devcloudpipeline
 *
 * devcloud pipeline api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 流水线参数详情
type Workflow struct {
	Parameter *PipelineParam `json:"parameter"`
	Source    *Source        `json:"source"`
	// 流水线名字
	Name string `json:"name"`
	// 项目ID
	ProjectId string `json:"project_id"`
	// 项目名字
	ProjectName string `json:"project_name"`
}

func (o Workflow) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"Workflow", string(data)}, " ")
}
