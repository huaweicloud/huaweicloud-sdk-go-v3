package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTagDto 自定义标签键值对。
type DeleteTagDto struct {

	// 标签的键
	Key string `json:"key"`

	// 标签的值，可以为空字符串，但不能为null
	Value *string `json:"value,omitempty"`
}

func (o DeleteTagDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTagDto struct{}"
	}

	return strings.Join([]string{"DeleteTagDto", string(data)}, " ")
}
