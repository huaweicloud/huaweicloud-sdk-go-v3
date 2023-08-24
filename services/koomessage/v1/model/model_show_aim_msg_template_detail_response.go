package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAimMsgTemplateDetailResponse Response Object
type ShowAimMsgTemplateDetailResponse struct {

	// 模板ID。
	TemplateId *string `json:"template_id,omitempty"`

	// 模板名称。
	TemplateName *string `json:"template_name,omitempty"`

	// 创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 模板类型。
	TemplateType *string `json:"template_type,omitempty"`

	// 应用ID。
	AppId *string `json:"app_id,omitempty"`

	// 应用名称。
	AppName *string `json:"app_name,omitempty"`

	// 签名ID。
	SignatureId *string `json:"signature_id,omitempty"`

	// 模板内容。
	TemplateContent *string `json:"template_content,omitempty"`

	// 申请描述。
	TemplateDesc *string `json:"template_desc,omitempty"`

	// 是否有变量。
	HasVariable *string `json:"has_variable,omitempty"`

	// 流程状态。
	FlowStatus *string `json:"flow_status,omitempty"`

	// 模板状态。
	Status *string `json:"status,omitempty"`

	// 是否是通用模板。
	UniversalTemplate *int32 `json:"universal_template,omitempty"`

	// 催审状态。
	UrgeStatus *string `json:"urge_status,omitempty"`

	// 催审时间。
	UrgeTime *string `json:"urge_time,omitempty"`

	// 催审描述。
	UrgeDesc *string `json:"urge_desc,omitempty"`

	// 审批描述。
	ReviewDesc     *string `json:"review_desc,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAimMsgTemplateDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAimMsgTemplateDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowAimMsgTemplateDetailResponse", string(data)}, " ")
}
