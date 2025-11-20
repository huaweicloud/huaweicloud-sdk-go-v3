package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ParaGroupUpdate struct {

	// 名称。
	Name *string `json:"name,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 参数值。
	Values *interface{} `json:"values,omitempty"`
}

func (o ParaGroupUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParaGroupUpdate struct{}"
	}

	return strings.Join([]string{"ParaGroupUpdate", string(data)}, " ")
}
