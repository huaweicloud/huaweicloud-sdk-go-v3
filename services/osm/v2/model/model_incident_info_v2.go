package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IncidentInfoV2 struct {

	// 子用户id
	XcustomerId *string `json:"xcustomer_id,omitempty"`

	// 子用户名称
	XcustomerName *string `json:"xcustomer_name,omitempty"`

	// 状态 0：待受理 1：处理中 2：待确认结果 3：已完成 4：已撤销 12：无效 17： 待反馈
	Status int32 `json:"status"`

	// 工单id
	IncidentId string `json:"incident_id"`

	// 问题类型名称
	BusinessTypeName string `json:"business_type_name"`

	// 工单类型名称
	IncidentTypeName string `json:"incident_type_name"`

	// 产品类型名称
	ProductCategoryName *string `json:"product_category_name,omitempty"`

	// 服务类型
	IncidentServiceType *string `json:"incident_service_type,omitempty"`

	// 客户id
	CustomerId string `json:"customer_id"`

	// 区域名称
	DcName string `json:"dc_name"`

	// 简要描述
	SimpleDescription string `json:"simple_description"`

	// 问题归属方
	RootCause *string `json:"root_cause,omitempty"`

	// 解决方案
	Resolution *string `json:"resolution,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time"`

	// 解决时间
	ConfirmedTime *sdktime.SdkTime `json:"confirmed_time,omitempty"`

	// 标签列表
	LabelList *[]LabelInfo `json:"label_list,omitempty"`
}

func (o IncidentInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncidentInfoV2 struct{}"
	}

	return strings.Join([]string{"IncidentInfoV2", string(data)}, " ")
}
