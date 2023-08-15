package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataobjectSearchCondition search condition
type DataobjectSearchCondition struct {

	// conditions
	Conditions *[]ConditonInfo `json:"conditions,omitempty"`

	// conditions
	Logics *[]string `json:"logics,omitempty"`
}

func (o DataobjectSearchCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataobjectSearchCondition struct{}"
	}

	return strings.Join([]string{"DataobjectSearchCondition", string(data)}, " ")
}
