package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LabelSelector struct {
	MatchLabels map[string]string `json:"matchLabels,omitempty"`

	MatchExpressions *[]FieldSelector `json:"matchExpressions,omitempty"`
}

func (o LabelSelector) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LabelSelector struct{}"
	}

	return strings.Join([]string{"LabelSelector", string(data)}, " ")
}
