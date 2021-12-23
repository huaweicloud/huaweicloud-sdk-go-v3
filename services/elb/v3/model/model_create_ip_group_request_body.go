package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type CreateIpGroupRequestBody struct {
	Ipgroup *CreateIpGroupOption `json:"ipgroup"`
}

func (o CreateIpGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpGroupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateIpGroupRequestBody", string(data)}, " ")
}
