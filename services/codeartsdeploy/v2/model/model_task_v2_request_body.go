package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskV2RequestBody struct {

	// 模板id
	TemplateId *string `json:"template_id,omitempty"`

	// 部署编排列表信息
	OperationList *[]DeployV2OperationsDo `json:"operation_list,omitempty"`
}

func (o TaskV2RequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskV2RequestBody struct{}"
	}

	return strings.Join([]string{"TaskV2RequestBody", string(data)}, " ")
}
