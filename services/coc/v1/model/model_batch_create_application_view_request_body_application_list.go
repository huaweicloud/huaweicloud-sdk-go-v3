package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateApplicationViewRequestBodyApplicationList struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 父节点名称
	ParentName *string `json:"parent_name,omitempty"`

	// 层级，从1开始
	Level *string `json:"level,omitempty"`
}

func (o BatchCreateApplicationViewRequestBodyApplicationList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateApplicationViewRequestBodyApplicationList struct{}"
	}

	return strings.Join([]string{"BatchCreateApplicationViewRequestBodyApplicationList", string(data)}, " ")
}
