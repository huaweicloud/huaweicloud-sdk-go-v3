package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceTag 专属主机标签结构体
type ResourceTag struct {

	// 键。 长度不超过36个Unicode字符。 不能为空。 不能包含以下ASCII非打印字符：“=”,“*”,“<”,“>”,“\\”,“|”,“/”,“,”。
	Key string `json:"key"`

	// 值。 长度不超过43个Unicode字符。 不能包含以下ASCII非打印字符：“=”,“*”,“<”,“>”,“\\”,“|”,“/”,“,”。
	Value string `json:"value"`
}

func (o ResourceTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTag struct{}"
	}

	return strings.Join([]string{"ResourceTag", string(data)}, " ")
}
