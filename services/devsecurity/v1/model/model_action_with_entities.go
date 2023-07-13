package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ActionWithEntities struct {

	// 动作列表
	Actions *[]string `json:"actions,omitempty"`

	// 实体列表
	Entities *[]string `json:"entities,omitempty"`
}

func (o ActionWithEntities) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionWithEntities struct{}"
	}

	return strings.Join([]string{"ActionWithEntities", string(data)}, " ")
}
