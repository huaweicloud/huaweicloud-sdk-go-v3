package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 热线模型。
type TelModel struct {

	// 电话号码（只能包含数字和”-“，且开头和结尾必须为数字）。
	Tel string `json:"tel"`

	// 号码用途。  > 号码用途长度范围为1-30个字符，中文占2个字符，英文占1个字符。
	Usage string `json:"usage"`
}

func (o TelModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TelModel struct{}"
	}

	return strings.Join([]string{"TelModel", string(data)}, " ")
}
