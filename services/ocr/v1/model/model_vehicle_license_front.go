package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VehicleLicenseFront
type VehicleLicenseFront struct {

	// 号牌号码。
	Number *string `json:"number,omitempty"`

	// 车辆类型。
	VehicleType *string `json:"vehicle_type,omitempty"`

	// 所有人。
	Name *string `json:"name,omitempty"`

	// 住址。
	Address *string `json:"address,omitempty"`

	// 使用性质。
	UseCharacter *string `json:"use_character,omitempty"`

	// 品牌型号。
	Model *string `json:"model,omitempty"`

	// 车辆识别代号。
	Vin *string `json:"vin,omitempty"`

	// 发动机号码。
	EngineNo *string `json:"engine_no,omitempty"`

	// 注册日期。
	RegisterDate *string `json:"register_date,omitempty"`

	// 发证日期。
	IssueDate *string `json:"issue_date,omitempty"`

	// 发证机关。
	IssuingAuthority *string `json:"issuing_authority,omitempty"`

	// 文本框在原图位置。输出左上、右上、右下、左下四个点坐标。  当“return_text_location”设置为“true”时才返回。
	TextLocation *interface{} `json:"text_location,omitempty"`

	AlarmResult *VehicleLicenseAlarmResult `json:"alarm_result,omitempty"`

	AlarmConfidence *VehicleLicenseAlarmConfidence `json:"alarm_confidence,omitempty"`
}

func (o VehicleLicenseFront) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VehicleLicenseFront struct{}"
	}

	return strings.Join([]string{"VehicleLicenseFront", string(data)}, " ")
}
