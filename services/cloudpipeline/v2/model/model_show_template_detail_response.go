package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowTemplateDetailResponse struct {
	// 模板ID

	TemplateId *string `json:"template_id,omitempty"`
	// 模板名字

	TemplateName *string `json:"template_name,omitempty"`
	// 模板类型

	TemplateType *string `json:"template_type,omitempty"`
	// 用户ID

	UserId *string `json:"user_id,omitempty"`
	// 用户名字

	UserName *string `json:"user_name,omitempty"`
	// 租户ID

	DomainId *string `json:"domain_id,omitempty"`
	// 租户名字

	DomainName *string `json:"domain_name,omitempty"`
	// 是否内置模板

	IsBuildIn *bool `json:"is_build_in,omitempty"`
	// region

	Region *string `json:"region,omitempty"`
	// 项目ID

	ProjectId *string `json:"project_id,omitempty"`
	// 项目名字

	ProjectName *string `json:"project_name,omitempty"`
	// 创建时间

	CreateTime *string `json:"create_time,omitempty"`
	// 修改时间

	LastModifyTime *string `json:"last_modify_time,omitempty"`
	// 是否关注

	IsWatch *bool `json:"is_watch,omitempty"`
	// 模板描述

	Description *string `json:"description,omitempty"`
	// 模板参数

	Parameter *[]TemplateParam `json:"parameter,omitempty"`
	// 编排flow详情，描述流水线内各阶段任务的串并行关系。map类型数据，key为阶段名字，默认第一阶段initial，最后阶段为final，其余名字以'state_数字'标识。value为该阶段内任务(以'Task_数字'标识)以及后续阶段的标识。本字段为描述流水线基础编排数据之一，建议可通过流水线真实界面基于模板创建接口中获取

	Flow map[string]map[string]string `json:"flow,omitempty"`
	// 编排State详情，map类型数据。本字段为描述流水线基础编排数据之一，建议可通过流水线真实界面基于模板创建接口中获取

	States         map[string]interface{} `json:"states,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowTemplateDetailResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTemplateDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowTemplateDetailResponse", string(data)}, " ")
}
