package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DriverLicenseResult struct {

	// 驾驶证类型。 normal：纸质驾驶证 electronic：电子驾驶证
	Type *string `json:"type,omitempty" xml:"type"`

	// 驾驶证号。
	Number *string `json:"number,omitempty" xml:"number"`

	// 姓名。
	Name *string `json:"name,omitempty" xml:"name"`

	// 性别。
	Sex *string `json:"sex,omitempty" xml:"sex"`

	// 国籍。
	Nationality *string `json:"nationality,omitempty" xml:"nationality"`

	// 住址。
	Address *string `json:"address,omitempty" xml:"address"`

	// 出生日期。
	Birth *string `json:"birth,omitempty" xml:"birth"`

	// 初次领证日期。
	IssueDate *string `json:"issue_date,omitempty" xml:"issue_date"`

	// 准驾类型。
	Class *string `json:"class,omitempty" xml:"class"`

	// 有效起始日期。
	ValidFrom *string `json:"valid_from,omitempty" xml:"valid_from"`

	// 有效结束日期。
	ValidTo *string `json:"valid_to,omitempty" xml:"valid_to"`

	// 发证机关。
	IssuingAuthority *string `json:"issuing_authority,omitempty" xml:"issuing_authority"`

	// 档案编号。
	FileNumber *string `json:"file_number,omitempty" xml:"file_number"`

	// 记录。
	Record *string `json:"record,omitempty" xml:"record"`

	// 累积记分。
	AccumulatedScores *string `json:"accumulated_scores,omitempty" xml:"accumulated_scores"`

	// 状态。
	Status *[]string `json:"status,omitempty" xml:"status"`

	// 生成时间。
	GenerationDate *string `json:"generation_date,omitempty" xml:"generation_date"`

	// 当前时间。
	CurrentTime *string `json:"current_time,omitempty" xml:"current_time"`

	// 对应所有在原图上识别到的字段位置信息，包含所有文字区域四个顶点的二维坐标（x,y）。采用图像坐标系，坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	TextLocation *interface{} `json:"text_location,omitempty" xml:"text_location"`
}

func (o DriverLicenseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DriverLicenseResult struct{}"
	}

	return strings.Join([]string{"DriverLicenseResult", string(data)}, " ")
}
