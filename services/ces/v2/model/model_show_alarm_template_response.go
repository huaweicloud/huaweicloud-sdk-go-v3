package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlarmTemplateResponse Response Object
type ShowAlarmTemplateResponse struct {

	// 告警模板的ID，以at开头，后跟字母、数字，长度最长为64
	TemplateId *string `json:"template_id,omitempty"`

	// 告警模板的名称，以字母或汉字开头，可包含字母、数字、汉字、_、-，长度范围[1,128]
	TemplateName *string `json:"template_name,omitempty"`

	TemplateType *TemplateType `json:"template_type,omitempty"`

	// 告警模板的创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 告警模板的描述，长度范围[0,256]，该字段默认值为空字符串
	TemplateDescription *string `json:"template_description,omitempty"`

	// 告警模板关联的告警规则数目
	AssociationAlarmTotal *int32 `json:"association_alarm_total,omitempty"`

	// 告警模板策略列表
	Policies       *[]AlarmTemplatePolicies `json:"policies,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowAlarmTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlarmTemplateResponse struct{}"
	}

	return strings.Join([]string{"ShowAlarmTemplateResponse", string(data)}, " ")
}
