package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SampleParam
type SampleParam struct {

	// 训练集测试集划分方式： - TIME，时间比例 - RAMDOM，个数比例
	DivideType string `json:"divide_type"`

	// 训练数据占比。
	TrainRate *float64 `json:"train_rate,omitempty"`

	// 测试数据占比。
	TestRate *float64 `json:"test_rate,omitempty"`
}

func (o SampleParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SampleParam struct{}"
	}

	return strings.Join([]string{"SampleParam", string(data)}, " ")
}
