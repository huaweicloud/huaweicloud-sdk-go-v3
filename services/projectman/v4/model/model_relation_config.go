package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RelationConfig struct {

	// key为工作项类型，value为该类型的所有关联关系
	Relations map[string][]Relation `json:"relations,omitempty"`
}

func (o RelationConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelationConfig struct{}"
	}

	return strings.Join([]string{"RelationConfig", string(data)}, " ")
}
