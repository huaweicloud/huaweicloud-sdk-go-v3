package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BuildParameter struct {

	// 参数名
	Name *string `json:"name,omitempty"`

	// 参数值
	Value *string `json:"value,omitempty"`
}

func (o BuildParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BuildParameter struct{}"
	}

	return strings.Join([]string{"BuildParameter", string(data)}, " ")
}
