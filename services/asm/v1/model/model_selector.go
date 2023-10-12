package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Selector struct {
	FieldSelector *FieldSelector `json:"fieldSelector,omitempty"`

	LabelSelector *LabelSelector `json:"labelSelector,omitempty"`
}

func (o Selector) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Selector struct{}"
	}

	return strings.Join([]string{"Selector", string(data)}, " ")
}
