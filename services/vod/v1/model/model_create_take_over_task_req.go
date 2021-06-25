package model

import (
	"encoding/json"

	"strings"
)

type CreateTakeOverTaskReq struct {
	// 桶 <br/>

	Bucket string `json:"bucket"`
	// 目录

	Object string `json:"object"`
	// 指定托管文件后缀名 - 如果不传，或传空数组，则不会进行文件后缀名过滤。

	Suffix *[]string `json:"suffix,omitempty"`
	// 转码模板组 <br/> 若未设置template_group_name和workflow_name，则不进行转码动作。 若同时设置了“template_group_name”和“workflow_name”字段，则“template_group_name”字段生效。

	TemplateGroupName *string `json:"template_group_name,omitempty"`
	// '工作流ID' 若同时设置了“template_group_name”和“workflow_name”字段，则“template_group_name”字段生效。

	WorkflowName *string `json:"workflow_name,omitempty"`
	// 托管类型：0表示存储到点播桶 1表示存储在租户桶  2表示存储到租户桶，并且存储路径与源文件一致。<br/>

	HostType *int32 `json:"host_type,omitempty"`
	// 输出桶 <br/> 约束： - host_type为1时，必选

	OutputBucket *string `json:"output_bucket,omitempty"`
	// 输出路径 <br/> 约束： - host_type为1时，必选

	OutputPath *string `json:"output_path,omitempty"`
}

func (o CreateTakeOverTaskReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTakeOverTaskReq struct{}"
	}

	return strings.Join([]string{"CreateTakeOverTaskReq", string(data)}, " ")
}
