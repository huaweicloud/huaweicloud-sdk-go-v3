package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetricTableItem struct {
	Allocated *Allocated `json:"allocated,omitempty"`

	Capacity *Capacity `json:"capacity,omitempty"`
}

func (o MetricTableItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricTableItem struct{}"
	}

	return strings.Join([]string{"MetricTableItem", string(data)}, " ")
}
