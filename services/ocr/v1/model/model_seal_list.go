package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SealList
type SealList struct {

	// 印章类型，当前支持circle（圆形章）、ellipse（椭圆章）、rectangle（方形章）、triangle（三角章）、rhombus（菱形章）五种。
	Type *string `json:"type,omitempty"`

	// 提取的单个印章base64编码图片。
	SealImage *string `json:"seal_image,omitempty"`

	// 印章位置的置信度。
	Confidence *float32 `json:"confidence,omitempty"`

	// 印章位置，列表形式，包含印章区域四个顶点的二维坐标（x,y），坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	Location *[][]int32 `json:"location,omitempty"`

	// 印章文本块列表。
	WordsBlockList *[]SealWordsBlockList `json:"words_block_list,omitempty"`
}

func (o SealList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SealList struct{}"
	}

	return strings.Join([]string{"SealList", string(data)}, " ")
}
