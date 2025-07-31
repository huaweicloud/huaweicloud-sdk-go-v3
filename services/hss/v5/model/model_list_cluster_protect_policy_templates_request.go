package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterProtectPolicyTemplatesRequest Request Object
type ListClusterProtectPolicyTemplatesRequest struct {

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 每页显示数量
	Limit int32 `json:"limit"`

	// 偏移量：指定返回记录的开始位置
	Offset int32 `json:"offset"`

	// 模板名称
	TemplateName *string `json:"template_name,omitempty"`

	// 模板类型
	TemplateType *string `json:"template_type,omitempty"`

	// 策略模板应用资源类型，多个资源类型通过分号份隔连接
	TargetKind *string `json:"target_kind,omitempty"`

	// 标签
	Tag *string `json:"tag,omitempty"`

	// 推荐等级
	Level *string `json:"level,omitempty"`
}

func (o ListClusterProtectPolicyTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterProtectPolicyTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListClusterProtectPolicyTemplatesRequest", string(data)}, " ")
}
