package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RealEstateCertificateResult struct {

	// 填发单位。
	Issuer *string `json:"issuer,omitempty"`

	// 填发日期。
	IssueDate *string `json:"issue_date,omitempty"`

	// 不动产证编号。
	RealEstateCertificateNo *string `json:"real_estate_certificate_no,omitempty"`

	// 抵押印章个数。
	MortgageSeals *int32 `json:"mortgage_seals,omitempty"`

	// 注销的抵押印章个数。
	CanceledMortgageSeals *int32 `json:"canceled_mortgage_seals,omitempty"`

	// 房屋坐落。
	EstateLocation *string `json:"estate_location,omitempty"`

	// 总楼层数。
	TotalFloors *string `json:"total_floors,omitempty"`

	// 所在层。
	Floor *string `json:"floor,omitempty"`

	// 建成年份。
	YearBuilt *string `json:"year_built,omitempty"`

	// 结构。
	Structure *string `json:"structure,omitempty"`

	// 建筑面积。
	Area *string `json:"area,omitempty"`

	// 印花税票个数。
	RevenueStamps *int32 `json:"revenue_stamps,omitempty"`

	// 产权证号。
	OwnershipCertificateNo *string `json:"ownership_certificate_no,omitempty"`

	// 房屋所有权人。
	EstateHolder *string `json:"estate_holder,omitempty"`

	// 权利人。
	Obligee *string `json:"obligee,omitempty"`

	// 共有情况。
	Ownership *string `json:"ownership,omitempty"`

	// 不动产单元号。
	PropertyUnitNo *string `json:"property_unit_no,omitempty"`

	// 权利类型。
	RightType *string `json:"right_type,omitempty"`

	// 权利性质。
	RightNature *string `json:"right_nature,omitempty"`

	// 使用用途。
	Usage *string `json:"usage,omitempty"`

	// 设计、规划用途。
	IntendedUsage *string `json:"intended_usage,omitempty"`

	// 各个字段的置信度。
	Confidence *interface{} `json:"confidence,omitempty"`
}

func (o RealEstateCertificateResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RealEstateCertificateResult struct{}"
	}

	return strings.Join([]string{"RealEstateCertificateResult", string(data)}, " ")
}
