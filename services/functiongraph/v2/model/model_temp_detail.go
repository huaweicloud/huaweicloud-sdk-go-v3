package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TempDetail struct {

	// 模板输入
	Input *string `json:"input,omitempty"`

	// 模板输出
	Output *string `json:"output,omitempty"`

	// 警告信息
	Warning *string `json:"warning,omitempty"`
}

func (o TempDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TempDetail struct{}"
	}

	return strings.Join([]string{"TempDetail", string(data)}, " ")
}
