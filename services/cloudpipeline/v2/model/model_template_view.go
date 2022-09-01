package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 流水线创建状态响应体
type TemplateView struct {

	// 模板ID
	TemplateId string `json:"template_id" xml:"template_id"`

	// 模板名字
	TemplateName string `json:"template_name" xml:"template_name"`

	// 模板类型
	TemplateType string `json:"template_type" xml:"template_type"`

	// 模板编辑URL
	TemplateUrl string `json:"template_url" xml:"template_url"`

	// 用户ID
	UserId string `json:"user_id" xml:"user_id"`

	// 用户名字
	UserName string `json:"user_name" xml:"user_name"`

	// 租户ID
	DomainId string `json:"domain_id" xml:"domain_id"`

	// 租户名字
	DomainName string `json:"domain_name" xml:"domain_name"`

	// 是否内置模板
	IsBuildIn bool `json:"is_build_in" xml:"is_build_in"`

	// 系统模板region为Cloud Pipeline。自定义模板region为实际region
	Region string `json:"region" xml:"region"`

	// 项目ID
	ProjectId string `json:"project_id" xml:"project_id"`

	// 项目名字
	ProjectName string `json:"project_name" xml:"project_name"`

	// 创建时间
	CreateTime string `json:"create_time" xml:"create_time"`

	// 修改时间
	LastModifyTime string `json:"last_modify_time" xml:"last_modify_time"`

	// 是否关注
	IsWatch bool `json:"is_watch" xml:"is_watch"`

	// 模板描述
	Description string `json:"description" xml:"description"`

	// 模板参数
	Parameter []TemplateParam `json:"parameter" xml:"parameter"`

	// 编排flow详情，描述流水线内各阶段任务的串并行关系。map类型数据，key为阶段名字，默认第一阶段initial，最后阶段为final，其余名字以'state_数字'标识。value为该阶段内任务(以'Task_数字'标识)以及后续阶段的标识。本字段为描述流水线基础编排数据之一，建议可通过流水线真实界面基于模板创建接口中获取
	Flow map[string]map[string]string `json:"flow" xml:"flow"`

	// 编排State详情，map类型数据。本字段为描述流水线基础编排数据之一，建议可通过流水线真实界面基于模板创建接口中获取
	States map[string]TemplateState `json:"states" xml:"states"`

	// 是否可以修改
	CanUpdate bool `json:"can_update" xml:"can_update"`

	// 是否可以删除
	CanDelete bool `json:"can_delete" xml:"can_delete"`

	// 是否需要代码仓库
	NeedHub bool `json:"need_hub" xml:"need_hub"`
}

func (o TemplateView) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateView struct{}"
	}

	return strings.Join([]string{"TemplateView", string(data)}, " ")
}
