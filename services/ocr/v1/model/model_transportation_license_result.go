package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

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

	VehicleWeight *string `json:"vehicle_weight,omitempty"`
	// 车辆尺寸。

	VehicleSize *string `json:"vehicle_size,omitempty"`
	// 核发机关（非必有，依赖对应运输证板式）。

	IssuingAuthority *string `json:"issuing_authority,omitempty"`
	// 签发日期（非必有，依赖对应运输证板式）。

	IssueDate *string `json:"issue_date,omitempty"`
	// 业户地址（非必有，依赖对应运输证板式）。

	OwnerAddress *string `json:"owner_address,omitempty"`
	// 经济类型（非必有，依赖对应运输证板式）。

	EconomicType *string `json:"economic_type,omitempty"`
	// 经营许可证号（非必有，依赖对应运输证板式）。

	BusinessCertificate *string `json:"business_certificate,omitempty"`
	// 道路普通货物运输（非必有，依赖对应运输证板式）。

	BusinessScope *string `json:"business_scope,omitempty"`
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
