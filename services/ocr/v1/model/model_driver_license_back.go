package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DriverLicenseBack struct {

	// 驾驶证类型。 normal：纸质驾驶证 electronic：电子驾驶证
	Type *string `json:"type,omitempty"`

	// 驾驶证号。
	Number *string `json:"number,omitempty"`

	// 姓名。
	Name *string `json:"name,omitempty"`

	// 发证机关。
	IssuingAuthority *string `json:"issuing_authority,omitempty"`

	// 住址。
	Address *string `json:"address,omitempty"`

	// 档案编号。 > 说明：当驾驶证类型为纸质驾驶证时才返回。
	FileNumber *string `json:"file_number,omitempty"`

	// 记录。
	Record *string `json:"record,omitempty"`

	// 对应所有在原图上识别到的字段位置信息，包含所有文字区域四个顶点的二维坐标（x,y）。采用图像坐标系，坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	TextLocation *interface{} `json:"text_location,omitempty"`
}

func (o DriverLicenseBack) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DriverLicenseBack struct{}"
	}

	return strings.Join([]string{"DriverLicenseBack", string(data)}, " ")
}
