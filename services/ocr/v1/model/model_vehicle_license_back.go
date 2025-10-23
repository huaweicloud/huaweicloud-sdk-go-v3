package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VehicleLicenseBack
type VehicleLicenseBack struct {

	// 号牌号码。
	Number *string `json:"number,omitempty"`

	// 档案编码。
	FileNo *string `json:"file_no,omitempty"`

	// 核定载人数。
	ApprovedPassengers *string `json:"approved_passengers,omitempty"`

	// 总质量。
	GrossMass *string `json:"gross_mass,omitempty"`

	// 整备质量。
	UnladenMass *string `json:"unladen_mass,omitempty"`

	// 核定载质量。
	ApprovedLoad *string `json:"approved_load,omitempty"`

	// 外廓尺寸。
	Dimension *string `json:"dimension,omitempty"`

	// 准牵引总质量。
	TractionMass *string `json:"traction_mass,omitempty"`

	// 备注。
	Remarks *string `json:"remarks,omitempty"`

	// 检验记录。
	InspectionRecord *string `json:"inspection_record,omitempty"`

	// 条码号。
	CodeNumber *string `json:"code_number,omitempty"`

	// 能源类型。
	EnergyType *string `json:"energy_type,omitempty"`

	// 文本框在原图位置。输出左上、右上、右下、左下四个点坐标。  当“return_text_location”设置为“true”时才返回。
	TextLocation *interface{} `json:"text_location,omitempty"`
}

func (o VehicleLicenseBack) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VehicleLicenseBack struct{}"
	}

	return strings.Join([]string{"VehicleLicenseBack", string(data)}, " ")
}
