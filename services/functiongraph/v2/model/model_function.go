package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Function 函数信息
type Function struct {

	// 函数名称，在单个流程中，名称需要唯一
	Name string `json:"name"`

	// 函数调用URN
	Operation string `json:"operation"`

	// 函数扩展属性，由用户自己定制
	Metadata *interface{} `json:"metadata,omitempty"`
}

func (o Function) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Function struct{}"
	}

	return strings.Join([]string{"Function", string(data)}, " ")
}
