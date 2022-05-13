package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type HealthCodeResult struct {

	// 姓名
	Name string `json:"name"`

	// 健康码更新时间
	Time string `json:"time"`

	// 健康码颜色，可选值包括： - \"green\" - \"yellow\" - \"red\" - \"gray\"
	Color string `json:"color"`

	// 各个字段的置信度
	Confidence *interface{} `json:"confidence"`
}

func (o HealthCodeResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthCodeResult struct{}"
	}

	return strings.Join([]string{"HealthCodeResult", string(data)}, " ")
}
