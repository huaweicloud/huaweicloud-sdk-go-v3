package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DriverLicenseResult struct {

	// 驾驶证类型。 normal：纸质驾驶证 electronic：电子驾驶证
	Type *string `json:"type,omitempty"`

	// 驾驶证号。
	Number *string `json:"number,omitempty"`

	// 姓名。
	Name *string `json:"name,omitempty"`

	// 性别。
	Sex *string `json:"sex,omitempty"`

	// 国籍。
	Nationality *string `json:"nationality,omitempty"`

	// 住址。
	Address *string `json:"address,omitempty"`

	// 出生日期。
	Birth *string `json:"birth,omitempty"`

	// 初次领证日期。
	IssueDate *string `json:"issue_date,omitempty"`

	// 准驾类型。
	Class *string `json:"class,omitempty"`

	// 有效起始日期。
	ValidFrom *string `json:"valid_from,omitempty"`

	// 有效结束日期。
	ValidTo *string `json:"valid_to,omitempty"`

	// 发证机关。
	IssuingAuthority *string `json:"issuing_authority,omitempty"`

	// 档案编号。
	FileNumber *string `json:"file_number,omitempty"`

	// 记录。
	Record *string `json:"record,omitempty"`

	// 累积记分。
	AccumulatedScores *string `json:"accumulated_scores,omitempty"`

	// 状态。
	Status *[]string `json:"status,omitempty"`

	// 生成时间。
	GenerationDate *string `json:"generation_date,omitempty"`

	// 当前时间。
	CurrentTime *string `json:"current_time,omitempty"`

	// 对应所有在原图上识别到的字段位置信息，包含所有文字区域四个顶点的二维坐标（x,y）。采用图像坐标系，坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	TextLocation *interface{} `json:"text_location,omitempty"`

	Front *DriverLicenseFront `json:"front,omitempty"`

	Back *DriverLicenseBack `json:"back,omitempty"`
}

func (o DriverLicenseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DriverLicenseResult struct{}"
	}

	return strings.Join([]string{"DriverLicenseResult", string(data)}, " ")
}
