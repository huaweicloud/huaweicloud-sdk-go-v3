package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskV2RequestBody struct {

	// 部署系统，deployTemplate：部署模板
	DeploySystem *string `json:"deploy_system,omitempty"`

	// 模板id
	TemplateId *string `json:"template_id,omitempty"`

	// 是否切换新系统支持，1：是 0：否
	AnsibleType *int32 `json:"ansible_type,omitempty"`

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
