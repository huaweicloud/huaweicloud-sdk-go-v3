package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FieldSelector struct {
	Key *string `json:"key,omitempty"`

	Operator *string `json:"operator,omitempty"`

	Values *[]string `json:"values,omitempty"`
}

func (o FieldSelector) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FieldSelector struct{}"
	}

	return strings.Join([]string{"FieldSelector", string(data)}, " ")
}
