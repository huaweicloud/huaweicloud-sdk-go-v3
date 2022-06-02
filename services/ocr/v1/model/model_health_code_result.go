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

	// 各个字段的置信度。
	Confidence *interface{} `json:"confidence"`

	// 代表检测识别出来的文字块数目。
	WordsBlockCount int32 `json:"words_block_count"`

	// 识别文字块列表，输出顺序从左到右，从上到下。
	WordsBlockList []HealthCodeWordsBlockList `json:"words_block_list"`
}

func (o HealthCodeResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthCodeResult struct{}"
	}

	return strings.Join([]string{"HealthCodeResult", string(data)}, " ")
}
