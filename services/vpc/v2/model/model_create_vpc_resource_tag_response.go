package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateVpcResourceTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateVpcResourceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcResourceTagResponse struct{}"
	}

	return strings.Join([]string{"CreateVpcResourceTagResponse", string(data)}, " ")
}
