package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 按标签过滤资源后返回结果的排序方式
type Sorted struct {

	// 按key值对请求内容进行排序
	Key *string `json:"key,omitempty" xml:"key"`

	// 是否采用倒序
	Reverse *bool `json:"reverse,omitempty" xml:"reverse"`
}

func (o Sorted) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Sorted struct{}"
	}

	return strings.Join([]string{"Sorted", string(data)}, " ")
}
