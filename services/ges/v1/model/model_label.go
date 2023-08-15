package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Label GES 元数据中lable类型
type Label struct {

	// Label名称。
	Name *string `json:"name,omitempty"`

	// 属性Map
	Properties *interface{} `json:"properties,omitempty"`
}

func (o Label) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Label struct{}"
	}

	return strings.Join([]string{"Label", string(data)}, " ")
}
