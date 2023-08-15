package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CpiTaskData CPI任务的请求体
type CpiTaskData struct {

	// 蛋白质FASTA标题
	Header *string `json:"header,omitempty"`

	// 蛋白质FASTA序列
	Fasta string `json:"fasta"`

	// 分子SMILES表达式列表
	SmilesList []string `json:"smiles_list"`

	// 打分阈值，分值必须大于该阈值才会返回
	Threshold *float32 `json:"threshold,omitempty"`

	// 期望最大返回条目数（排序后取Top）
	NumResults *int32 `json:"num_results,omitempty"`

	// 用户已开启的自定义属性集合
	CustomProps *[]CustomProp `json:"custom_props,omitempty"`
}

func (o CpiTaskData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CpiTaskData struct{}"
	}

	return strings.Join([]string{"CpiTaskData", string(data)}, " ")
}
