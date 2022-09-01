package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IncidentDetailInfoV2 struct {

	// 子用户id
	XcustomerId *string `json:"xcustomer_id,omitempty" xml:"xcustomer_id"`

	// 子用户名称
	XcustomerName *string `json:"xcustomer_name,omitempty" xml:"xcustomer_name"`

	// 状态 0：待受理 1：处理中 2：待确认结果 3：已完成 4：已撤销 12：无效 17： 待反馈
	Status int32 `json:"status" xml:"status"`

	// 评价内容
	Judgement string `json:"judgement" xml:"judgement"`

	// 工单id
	IncidentId *string `json:"incident_id,omitempty" xml:"incident_id"`

	// 问题类型名称
	BusinessTypeName string `json:"business_type_name" xml:"business_type_name"`

	// 工单类型名称
	IncidentTypeName string `json:"incident_type_name" xml:"incident_type_name"`

	// 客户id
	CustomerId string `json:"customer_id" xml:"customer_id"`

	// 区域名称
	DcName string `json:"dc_name" xml:"dc_name"`

	// 简要描述
	SimpleDescription string `json:"simple_description" xml:"simple_description"`

	// 来源名称
	SourceName string `json:"source_name" xml:"source_name"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time" xml:"create_time"`

	// 留言列表
	MessageList []IncidentMessageV2 `json:"message_list" xml:"message_list"`

	// 满意度列表
	IncidentSatisfaction []IncidentSatisfactionV2Do `json:"incident_satisfaction" xml:"incident_satisfaction"`

	// 是否授权
	IsAuthorized int32 `json:"is_authorized" xml:"is_authorized"`

	// 严重性名称
	SeverityName *string `json:"severity_name,omitempty" xml:"severity_name"`

	// 业务归属 0华为云 1BP伙伴 2ISV
	BusinessOwnership *int32 `json:"business_ownership,omitempty" xml:"business_ownership"`

	// 解决时间
	ResolveTime *int64 `json:"resolve_time,omitempty" xml:"resolve_time"`

	ExtInfo *IncidentDetailExtInfoV2 `json:"ext_info,omitempty" xml:"ext_info"`
}

func (o IncidentDetailInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncidentDetailInfoV2 struct{}"
	}

	return strings.Join([]string{"IncidentDetailInfoV2", string(data)}, " ")
}
