package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IncidentInfo struct {

	// 工单id
	IncidentId *string `json:"incidentId,omitempty"`

	// 类型
	Type *int32 `json:"type,omitempty"`

	// 业务分类名称
	BusinessTypeName *string `json:"businessTypeName,omitempty"`

	// 工单类型名称
	IncidentTypeName *string `json:"incidentTypeName,omitempty"`

	// 客户id
	CustomerId *string `json:"customerId,omitempty"`

	// 客户id
	XcustomerId *string `json:"xcustomerId,omitempty"`

	// 客户名称
	XcustomerName *string `json:"xcustomerName,omitempty"`

	// dc名称
	DcName *string `json:"dcName,omitempty"`

	// 简单说明
	SimpleDescribe *string `json:"simpleDescribe,omitempty"`

	// 工单状态
	Status *int32 `json:"status,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"createTime,omitempty"`

	// 是否灰度
	Gray *bool `json:"gray,omitempty"`

	// 标签列表
	LabelList *[]CaseQueryLabel `json:"labelList,omitempty"`
}

func (o IncidentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IncidentInfo struct{}"
	}

	return strings.Join([]string{"IncidentInfo", string(data)}, " ")
}
