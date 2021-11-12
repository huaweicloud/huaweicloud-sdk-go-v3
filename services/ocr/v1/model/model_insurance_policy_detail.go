package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InsurancePolicyDetail struct {
	// 对应识别出的文本内容

	Words *string `json:"words,omitempty"`
	// 对应识别出的四个顶点坐标

	Location *[][]int32 `json:"location,omitempty"`
}

func (o InsurancePolicyDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InsurancePolicyDetail struct{}"
	}

	return strings.Join([]string{"InsurancePolicyDetail", string(data)}, " ")
}
