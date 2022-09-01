package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IncidentInfoV2 struct {

	// 子用户id
	XcustomerId *string `json:"xcustomer_id,omitempty" xml:"xcustomer_id"`

	// 子用户名称
	XcustomerName *string `json:"xcustomer_name,omitempty" xml:"xcustomer_name"`

	// 状态 0：待受理 1：处理中 2：待确认结果 3：已完成 4：已撤销 12：无效 17： 待反馈
	Status int32 `json:"status" xml:"status"`

	// 工单id
	IncidentId string `json:"incident_id" xml:"incident_id"`

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

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time" xml:"create_time"`

	// 标签列表
	LabelList *[]LabelInfo `json:"label_list,omitempty" xml:"label_list"`
}

func (o IncidentInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncidentInfoV2 struct{}"
	}

	return strings.Join([]string{"IncidentInfoV2", string(data)}, " ")
}
