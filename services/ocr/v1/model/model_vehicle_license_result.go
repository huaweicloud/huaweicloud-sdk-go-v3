package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VehicleLicenseResult struct {

	// 号牌号码。
	Number *string `json:"number,omitempty" xml:"number"`

	// 车辆类型。
	VehicleType *string `json:"vehicle_type,omitempty" xml:"vehicle_type"`

	// 所有人。
	Name *string `json:"name,omitempty" xml:"name"`

	// 住址。
	Address *string `json:"address,omitempty" xml:"address"`

	// 使用性质。
	UseCharacter *string `json:"use_character,omitempty" xml:"use_character"`

	// 品牌型号。
	Model *string `json:"model,omitempty" xml:"model"`

	// 发动机号码。
	EngineNo *string `json:"engine_no,omitempty" xml:"engine_no"`

	// 车辆识别代号。
	Vin *string `json:"vin,omitempty" xml:"vin"`

	// 注册日期。
	RegisterDate *string `json:"register_date,omitempty" xml:"register_date"`

	// 发证日期。
	IssueDate *string `json:"issue_date,omitempty" xml:"issue_date"`

	// 发证机关。
	IssuingAuthority *string `json:"issuing_authority,omitempty" xml:"issuing_authority"`

	// 档案编码。
	FileNo *string `json:"file_no,omitempty" xml:"file_no"`

	// 核定载人数。
	ApprovedPassengers *string `json:"approved_passengers,omitempty" xml:"approved_passengers"`

	// 总质量。
	GrossMass *string `json:"gross_mass,omitempty" xml:"gross_mass"`

	// 整备质量。
	UnladenMass *string `json:"unladen_mass,omitempty" xml:"unladen_mass"`

	// 核定载质量。
	ApprovedLoad *string `json:"approved_load,omitempty" xml:"approved_load"`

	// 外廓尺寸。
	Dimension *string `json:"dimension,omitempty" xml:"dimension"`

	// 准牵引总质量。
	TractionMass *string `json:"traction_mass,omitempty" xml:"traction_mass"`

	// 备注。
	Remarks *string `json:"remarks,omitempty" xml:"remarks"`

	// 检验记录。
	InspectionRecord *string `json:"inspection_record,omitempty" xml:"inspection_record"`

	// 条码号。
	CodeNumber *string `json:"code_number,omitempty" xml:"code_number"`

	// 文本框在原图位置。输出左上、右上、右下、左下四个点坐标。  当“return_text_location”设置为“true”时才返回。
	TextLocation *interface{} `json:"text_location,omitempty" xml:"text_location"`
}

func (o VehicleLicenseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VehicleLicenseResult struct{}"
	}

	return strings.Join([]string{"VehicleLicenseResult", string(data)}, " ")
}
