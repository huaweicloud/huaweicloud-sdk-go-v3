package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TransportationLicenseResult struct {

	// 业户名称。
	OwnerName *string `json:"owner_name,omitempty"`

	// 道路运输证号。
	LicenseNumber *string `json:"license_number,omitempty"`

	// 车辆号牌。
	VehicleNumber *string `json:"vehicle_number,omitempty"`

	// 车辆类型。
	VehicleType *string `json:"vehicle_type,omitempty"`

	// 吨(座)位。
	MaximumCapacity *string `json:"maximum_capacity,omitempty"`

	// 车辆尺寸。
	VehicleSize *string `json:"vehicle_size,omitempty"`

	// 核发机关。
	IssuingAuthority *string `json:"issuing_authority,omitempty"`

	// 发证日期。
	IssueDate *string `json:"issue_date,omitempty"`

	// 业户地址。
	OwnerAddress *string `json:"owner_address,omitempty"`

	// 经济类型。
	EconomicType *string `json:"economic_type,omitempty"`

	// 经营许可证号。
	BusinessCertificate *string `json:"business_certificate,omitempty"`

	// 经营范围。
	BusinessScope *string `json:"business_scope,omitempty"`

	// 有效期。
	ExpiryDate *string `json:"expiry_date,omitempty"`

	// 审验有效期。
	ReviewExpiryDate *string `json:"review_expiry_date,omitempty"`

	// 技术等级评定。
	AssessedTechnicalLevel *string `json:"assessed_technical_level,omitempty"`

	// 相关字段的置信度信息，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。 置信度由算法给出，不直接等价于对应字段的准确率。
	Confidence *interface{} `json:"confidence,omitempty"`
}

func (o TransportationLicenseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransportationLicenseResult struct{}"
	}

	return strings.Join([]string{"TransportationLicenseResult", string(data)}, " ")
}
