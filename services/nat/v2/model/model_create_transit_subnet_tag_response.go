package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTransitSubnetTagResponse Response Object
type CreateTransitSubnetTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateTransitSubnetTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransitSubnetTagResponse struct{}"
	}

	return strings.Join([]string{"CreateTransitSubnetTagResponse", string(data)}, " ")
}
