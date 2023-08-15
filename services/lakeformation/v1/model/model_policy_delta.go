package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyDelta struct {
	Policy *Policy `json:"policy,omitempty"`
}

func (o PolicyDelta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyDelta struct{}"
	}

	return strings.Join([]string{"PolicyDelta", string(data)}, " ")
}
