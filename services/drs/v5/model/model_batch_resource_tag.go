package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchResourceTag 标签信息体
type BatchResourceTag struct {

	// 标签键。 - 长度为1-128个unicode字符。 - 可以包含任何语种字母、数字、空格和_.:=+-@，但首尾不能含有空格，不能以_sys_开头。
	Key string `json:"key"`

	// 标签值，删除标签时非必填。 - 最大长度255个unicode字符。 - 可以包含任何语种字母、数字、空格和_.:=+-@。
	Value *string `json:"value,omitempty"`
}

func (o BatchResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResourceTag struct{}"
	}

	return strings.Join([]string{"BatchResourceTag", string(data)}, " ")
}
