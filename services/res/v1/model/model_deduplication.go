package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 属性去重复，只能选择整型和字符串类型属性
type Deduplication struct {

	// 属性。
	Attributes *[]string `json:"attributes,omitempty"`
}

func (o Deduplication) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Deduplication struct{}"
	}

	return strings.Join([]string{"Deduplication", string(data)}, " ")
}
