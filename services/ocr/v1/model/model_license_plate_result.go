package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type LicensePlateResult struct {

	// 车牌内容。
	PlateNumber string `json:"plate_number"`

	// 当前版本支持的车牌底色类型：  - blue: 蓝色  - green: 绿色  - black: 黑色  - white: 白色  - yellow: 黄色
	PlateColor string `json:"plate_color"`

	// 车牌的区域位置信息，列表形式，包含文字区域四个顶点的二维坐标（x,y）;坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	PlateLocation [][]int32 `json:"plate_location"`

	// 相关字段的置信度信息，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。 置信度由算法给出，不直接等价于对应字段的准确率。
	Confidence float32 `json:"confidence"`
}

func (o LicensePlateResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LicensePlateResult struct{}"
	}

	return strings.Join([]string{"LicensePlateResult", string(data)}, " ")
}
