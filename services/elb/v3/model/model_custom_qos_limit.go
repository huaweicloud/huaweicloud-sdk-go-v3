package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomQosLimit struct {
	L4 *L4Limit `json:"l4,omitempty"`

	L7 *L7Limit `json:"l7,omitempty"`
}

func (o CustomQosLimit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomQosLimit struct{}"
	}

	return strings.Join([]string{"CustomQosLimit", string(data)}, " ")
}
