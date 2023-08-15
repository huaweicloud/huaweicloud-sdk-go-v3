package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTransitIpTagResponse Response Object
type CreateTransitIpTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateTransitIpTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransitIpTagResponse struct{}"
	}

	return strings.Join([]string{"CreateTransitIpTagResponse", string(data)}, " ")
}
