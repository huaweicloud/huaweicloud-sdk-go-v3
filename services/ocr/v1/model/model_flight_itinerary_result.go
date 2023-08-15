package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlightItineraryResult
type FlightItineraryResult struct {

	// 印刷序号。
	SerialNumber *string `json:"serial_number,omitempty"`

	// 旅客姓名。
	PassengerName *string `json:"passenger_name,omitempty"`

	// 有效身份证件号码。
	IdNumber *string `json:"id_number,omitempty"`

	// 备注。
	EndorsementsRestrictions *string `json:"endorsements_restrictions,omitempty"`

	// 订单号。
	OrderNumber *string `json:"order_number,omitempty"`

	// 票价。
	Fare *string `json:"fare,omitempty"`

	// 民航（CAAC)发展基金。
	CaacDevelopmentFund *string `json:"caac_development_fund,omitempty"`

	// 燃油附加费。
	FuelSurcharge *string `json:"fuel_surcharge,omitempty"`

	// 其他税费。
	OtherTaxes *string `json:"other_taxes,omitempty"`

	// 合计。
	Total *string `json:"total,omitempty"`

	// 电子客票号码。
	ETicketNumber *string `json:"e_ticket_number,omitempty"`

	// 验证码。
	CheckCode *string `json:"check_code,omitempty"`

	// 提示信息。
	ReferenceInformation *string `json:"reference_information,omitempty"`

	// 保险费。
	Insurance *string `json:"insurance,omitempty"`

	// 销售单位代号。
	AgentCode *string `json:"agent_code,omitempty"`

	// 填开单位。
	IssueOrganization *string `json:"issue_organization,omitempty"`

	// 填开日期。
	IssueDate *string `json:"issue_date,omitempty"`

	// 机票行程列表。
	ItineraryList *[]ItineraryList `json:"itinerary_list,omitempty"`

	// 相关字段的置信度信息，取值范围0~1。  置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。 置信度由算法给出，不直接等价于对应字段的准确率。  > 说明：  - （1）置信度中的相关字段均与返回值中的相关字段一一对应；  - （2）置信度中的itinerary_list的顺序与返回值中的itinerary_list的顺序是一致的。
	Confidence *interface{} `json:"confidence,omitempty"`
}

func (o FlightItineraryResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlightItineraryResult struct{}"
	}

	return strings.Join([]string{"FlightItineraryResult", string(data)}, " ")
}
