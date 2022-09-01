package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAuthorizationDetailResponse struct {

	// 授权id
	Id *int64 `json:"id,omitempty" xml:"id"`

	// 授权状态
	Status *int32 `json:"status,omitempty" xml:"status"`

	// 工单id
	IncidentId *string `json:"incident_id,omitempty" xml:"incident_id"`

	// 简要描述
	SimpleDescription *string `json:"simple_description,omitempty" xml:"simple_description"`

	// 授权资源类型id
	ResourceTypeId *string `json:"resource_type_id,omitempty" xml:"resource_type_id"`

	// 授权资源类型名称
	ResourceTypeName *string `json:"resource_type_name,omitempty" xml:"resource_type_name"`

	// 授权访问类型id
	VisitTypeId *string `json:"visit_type_id,omitempty" xml:"visit_type_id"`

	// 授权访问类型名称
	VisitTypeName *string `json:"visit_type_name,omitempty" xml:"visit_type_name"`

	// 授权生效时间
	AuthEffectiveTime *sdktime.SdkTime `json:"auth_effective_time,omitempty" xml:"auth_effective_time"`

	// 授权到期时间
	AuthExpireTime *sdktime.SdkTime `json:"auth_expire_time,omitempty" xml:"auth_expire_time"`

	// 拒绝原因
	RejectReason *string `json:"reject_reason,omitempty" xml:"reject_reason"`

	// 授权详情列表
	IncidentAuthDetailList *[]IncidentOrderAuthDetailInfoV2 `json:"incident_auth_detail_list,omitempty" xml:"incident_auth_detail_list"`

	// 子账号名称
	XcustomerName *string `json:"xcustomer_name,omitempty" xml:"xcustomer_name"`

	// 授权处理人名称
	AuthHandlerName *string `json:"auth_handler_name,omitempty" xml:"auth_handler_name"`

	// 委托名称
	AgencyName *string `json:"agency_name,omitempty" xml:"agency_name"`

	// 授权描述
	AuthDescribe *string `json:"auth_describe,omitempty" xml:"auth_describe"`

	// 授权内容Id
	ContentTypeId *string `json:"content_type_id,omitempty" xml:"content_type_id"`

	// 授权内容名称
	ContentTypeName *string `json:"content_type_name,omitempty" xml:"content_type_name"`
	HttpStatusCode  int     `json:"-"`
}

func (o ShowAuthorizationDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuthorizationDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowAuthorizationDetailResponse", string(data)}, " ")
}
