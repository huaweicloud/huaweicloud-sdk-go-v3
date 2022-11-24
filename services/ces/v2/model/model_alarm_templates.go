package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmTemplates struct {

	// 告警模板的ID，以at开头，后跟字母、数字，长度最长为64
	TemplateId string `json:"template_id"`

	// 告警模板的名称，以字母或汉字开头，可包含字母、数字、汉字、_、-，长度范围[1,128]
	TemplateName string `json:"template_name"`

	TemplateType *TemplateType `json:"template_type"`

	// 告警模板的创建时间
	CreateTime *sdktime.SdkTime `json:"create_time"`

	// 告警模板的描述，长度范围[0,256]，该字段默认值为空字符串
	TemplateDescription string `json:"template_description"`

	// 告警模板关联的告警规则数目
	AssociationAlarmTotal *int32 `json:"association_alarm_total,omitempty"`

	// 告警模板的告警策略总数
	PolicyTotal int32 `json:"policy_total"`

	// 服务列表告警策略数目统计
	PolicyStatistics []PolicyStatistics `json:"policy_statistics"`

	// 关联的资源分组
	AssociationResourceGroups *[]AssociationResourceGroup `json:"association_resource_groups,omitempty"`
}

func (o AlarmTemplates) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmTemplates struct{}"
	}

	return strings.Join([]string{"AlarmTemplates", string(data)}, " ")
}
