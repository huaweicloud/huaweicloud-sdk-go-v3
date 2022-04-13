package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type BaremetalServerTag struct {
	// 键。  - 最大长度127个unicode字符。  - key不能为空。

	Key string `json:"key"`
	// 值列表。  - 最多10个value。  - value不允许重复。  - 每个值最大长度255个unicode字符。  - 如果values为空则表示any_value。  - value之间为或的关系。

	Value string `json:"value"`
}

func (o BaremetalServerTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaremetalServerTag struct{}"
	}

	return strings.Join([]string{"BaremetalServerTag", string(data)}, " ")
}
