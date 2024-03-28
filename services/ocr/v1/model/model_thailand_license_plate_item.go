package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ThailandLicensePlateItem struct {

	// 车牌内容。
	PlateNumber *string `json:"plate_number,omitempty"`

	// 车牌的区域位置信息，列表形式，包含文字区域四个顶点的二维坐标（x,y）;坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	PlateLocation *[][]int32 `json:"plate_location,omitempty"`

	// 相关字段的置信度信息，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。注：置信度由算法给出，不直接等价于对应字段的准确率。
	Confidence *float32 `json:"confidence,omitempty"`

	// 车牌所属府
	Province *string `json:"province,omitempty"`
}

func (o ThailandLicensePlateItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ThailandLicensePlateItem struct{}"
	}

	return strings.Join([]string{"ThailandLicensePlateItem", string(data)}, " ")
}
