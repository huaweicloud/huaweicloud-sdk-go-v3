package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type WebImageWordsBlockList struct {
	// 文字块识别结果。

	Words *string `json:"words,omitempty"`
	// 相关字段的置信度信息，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。 置信度由算法给出，不直接等价于对应字段的准确率。

	Confidence *float32 `json:"confidence,omitempty"`
	// 文字块的区域位置信息，列表形式，包含文字区域四个顶点的二维坐标（x,y）;坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。

	Location *[][]int32 `json:"location,omitempty"`
	// 提取出的结构化JSON结果，该字典内的key值与入参列表extract_type的值一致，目前仅支持联系人信息提取，亦即key值为\"contact_info\"的字段。 若入参extract_type为空列表或该字段缺失时，不进行提取，此字段为空。

	ExtractedData *interface{} `json:"extracted_data,omitempty"`
	// 该字段表示提取的联系人信息，包括：姓名、联系电话、省市区以及详细地址。 若入参extract_type列表中无该字段，则此字段不存在。

	ContactInfo *interface{} `json:"contact_info,omitempty"`
	// 该字段表示返回图片宽高信息。 如入参extract_type列表中无该字段，则此字段不存在。

	ImageSize *interface{} `json:"image_size,omitempty"`
	// 传入image_size时的返回，为图像高度。

	Height *int32 `json:"height,omitempty"`
	// 传入image_size时的返回，为图像宽度。

	Width *int32 `json:"width,omitempty"`
	// 传入contact_info时的返回，为姓名。

	Name *string `json:"name,omitempty"`
	// 传入contact_info时的返回，联系电话。

	Phone *string `json:"phone,omitempty"`
	// 传入contact_info时的返回，省。

	Province *string `json:"province,omitempty"`
	// 传入contact_info时的返回，市。

	City *string `json:"city,omitempty"`
	// 传入contact_info时的返回，县区。

	District *string `json:"district,omitempty"`
	// 传入contact_info时的返回，详细地址（不含省市区）。

	DetailAddress *string `json:"detail_address,omitempty"`
}

func (o WebImageWordsBlockList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebImageWordsBlockList struct{}"
	}

	return strings.Join([]string{"WebImageWordsBlockList", string(data)}, " ")
}
