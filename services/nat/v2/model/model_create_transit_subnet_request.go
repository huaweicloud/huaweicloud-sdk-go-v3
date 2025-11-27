package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTransitSubnetRequest Request Object
type CreateTransitSubnetRequest struct {
	Body *CreateTransitSubnetRequestBody `json:"body,omitempty"`
}

func (o CreateTransitSubnetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransitSubnetRequest struct{}"
	}

	return strings.Join([]string{"CreateTransitSubnetRequest", string(data)}, " ")
}
