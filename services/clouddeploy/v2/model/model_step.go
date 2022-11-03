package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 执行步骤
type Step struct {

	// id
	Id *string `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 参数
	Params map[string]string `json:"params,omitempty"`

	// 是否开启
	Enable *bool `json:"enable,omitempty"`
}

func (o Step) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Step struct{}"
	}

	return strings.Join([]string{"Step", string(data)}, " ")
}
