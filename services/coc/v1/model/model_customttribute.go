package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Customttribute 当前执行脚本的自定义属性，仅用于表示本次执行的产生工单的信息
type Customttribute struct {

	// 自定义属性key
	Key string `json:"key"`

	// 自定义属性的value
	Value string `json:"value"`
}

func (o Customttribute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Customttribute struct{}"
	}

	return strings.Join([]string{"Customttribute", string(data)}, " ")
}
