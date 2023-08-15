package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IncidentOrderAuthV2 struct {

	// 授权id
	Id *string `json:"id,omitempty"`

	// 授权状态
	Status *int32 `json:"status,omitempty"`

	// 工单id
	IncidentId *string `json:"incident_id,omitempty"`

	// 简要描述
	SimpleDescription *string `json:"simple_description,omitempty"`

	// 授权描述
	AuthDescribe *string `json:"auth_describe,omitempty"`

	// 授权资源类型id
	ResourceTypeId *string `json:"resource_type_id,omitempty"`

	// 授权资源类型名称
	ResourceTypeName *string `json:"resource_type_name,omitempty"`

	// 授权访问类型id
	VisitTypeId *string `json:"visit_type_id,omitempty"`

	// 授权访问类型名称
	VisitTypeName *string `json:"visit_type_name,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 授权生效时间
	AuthEffectiveTime *sdktime.SdkTime `json:"auth_effective_time,omitempty"`

	// 授权到期时间
	AuthExpireTime *sdktime.SdkTime `json:"auth_expire_time,omitempty"`

	// 拒绝原因
	RejectReason *string `json:"reject_reason,omitempty"`

	// 主账号id
	CustomerId *string `json:"customer_id,omitempty"`

	// 子用户id
	XCustomerId *string `json:"x_customer_id,omitempty"`

	// 子用户名称
	XcustomerName *string `json:"xcustomer_name,omitempty"`

	// 承诺书内容
	ContentOfAuthLetter *string `json:"content_of_auth_letter,omitempty"`
}

func (o IncidentOrderAuthV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncidentOrderAuthV2 struct{}"
	}

	return strings.Join([]string{"IncidentOrderAuthV2", string(data)}, " ")
}
