package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProgressDetailV2 struct {

	// 进度百分比
	Ratio *string `json:"ratio,omitempty" xml:"ratio"`

	// 中文信息
	Info *string `json:"info,omitempty" xml:"info"`
}

func (o ProgressDetailV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProgressDetailV2 struct{}"
	}

	return strings.Join([]string{"ProgressDetailV2", string(data)}, " ")
}
