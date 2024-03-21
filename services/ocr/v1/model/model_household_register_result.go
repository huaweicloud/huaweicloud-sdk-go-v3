package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HouseholdRegisterResult struct {

	// 类型。参数为“首页”或“登记页”。
	Type *string `json:"type,omitempty"`

	// 户口本证件位置信息，列表形式，包含证件位置四个顶点的二维坐标（x,y）;坐标原点为图片左上角，x轴沿水平方向，y轴沿竖直方向。
	Location *[][]int32 `json:"location,omitempty"`

	Content *HouseholdRegisterContent `json:"content,omitempty"`

	// content中各个字段的置信度，取值范围0~1。置信度越大，本次识别的字段的可靠性越高，在统计意义上，置信度越大，准确率越高。置信度由算法给出，不直接等价于字段的准确率。
	Confidence map[string]float32 `json:"confidence,omitempty"`
}

func (o HouseholdRegisterResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HouseholdRegisterResult struct{}"
	}

	return strings.Join([]string{"HouseholdRegisterResult", string(data)}, " ")
}
