package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyCreateReq struct {
	Policy *PolicyCreate `json:"policy"`
}

func (o PolicyCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyCreateReq struct{}"
	}

	return strings.Join([]string{"PolicyCreateReq", string(data)}, " ")
}
