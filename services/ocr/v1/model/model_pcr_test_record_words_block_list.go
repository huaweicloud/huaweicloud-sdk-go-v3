package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PcrTestRecordWordsBlockList
type PcrTestRecordWordsBlockList struct {

	// 文字块识别结果。
	Words *string `json:"words,omitempty"`

	// 识别到的文字块的区域位置信息，列表形式，分别表示文字块4个顶点的（x,y）坐标；采用图像坐标系，图像坐标原点为图像左上角，x轴沿水平方向，y轴沿竖直方向。
	Location *[][]int32 `json:"location,omitempty"`

	// 字段的平均置信度，置信度越大，表示本次识别的对应字段的可靠性越高，在统计意义上，置信度越大，准确率越高。置信度由算法给出，不直接等价于对应字段的准确率。
	Confidence *float32 `json:"confidence,omitempty"`
}

func (o PcrTestRecordWordsBlockList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PcrTestRecordWordsBlockList struct{}"
	}

	return strings.Join([]string{"PcrTestRecordWordsBlockList", string(data)}, " ")
}
