package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Condition struct {
	PreNodeName *string `json:"preNodeName,omitempty" xml:"preNodeName"`

	Expression *string `json:"expression,omitempty" xml:"expression"`
}

func (o Condition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Condition struct{}"
	}

	return strings.Join([]string{"Condition", string(data)}, " ")
}
