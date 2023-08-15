package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRelation Create Relation
type CreateRelation struct {

	// search ids
	Ids *[]string `json:"ids,omitempty"`
}

func (o CreateRelation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRelation struct{}"
	}

	return strings.Join([]string{"CreateRelation", string(data)}, " ")
}
