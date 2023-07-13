package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceGroupTagRelation struct {

	// 键
	Key string `json:"key"`

	// 值
	Value *string `json:"value,omitempty"`
}

func (o ResourceGroupTagRelation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceGroupTagRelation struct{}"
	}

	return strings.Join([]string{"ResourceGroupTagRelation", string(data)}, " ")
}
