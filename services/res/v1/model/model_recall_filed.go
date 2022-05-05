package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 召回字段
type RecallFiled struct {

	// 字段名称。
	Name *string `json:"name,omitempty"`

	// 使用字段值的个数。
	Value *int32 `json:"value,omitempty"`
}

func (o RecallFiled) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecallFiled struct{}"
	}

	return strings.Join([]string{"RecallFiled", string(data)}, " ")
}
