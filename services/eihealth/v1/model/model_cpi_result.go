package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CPI任务的返回结果
type CpiResult struct {

	// 蛋白质FASTA标题
	Header *string `json:"header,omitempty"`

	// 蛋白质FASTA序列
	Fasta string `json:"fasta"`

	// 分子ADMET属性名列表
	PropNames *[]string `json:"prop_names,omitempty"`

	// 返回CPI的模型结果
	Result []CpiResultItem `json:"result"`
}

func (o CpiResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CpiResult struct{}"
	}

	return strings.Join([]string{"CpiResult", string(data)}, " ")
}
