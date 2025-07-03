package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CocTaskDetailV2 struct {

	// 节点类型
	Type *string `json:"type,omitempty"`

	// 节点KEY
	Key *string `json:"key,omitempty"`

	// 节点名称
	Name *string `json:"name,omitempty"`

	// 节点类型
	State *string `json:"state,omitempty"`

	// 操作列表
	Operations *[]CocTaskOperationDetailV2 `json:"operations,omitempty"`
}

func (o CocTaskDetailV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CocTaskDetailV2 struct{}"
	}

	return strings.Join([]string{"CocTaskDetailV2", string(data)}, " ")
}
