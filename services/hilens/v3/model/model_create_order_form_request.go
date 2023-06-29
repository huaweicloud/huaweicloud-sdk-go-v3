package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrderFormRequest Request Object
type CreateOrderFormRequest struct {
	Body *CreateSkillOrderFrom `json:"body,omitempty"`
}

func (o CreateOrderFormRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrderFormRequest struct{}"
	}

	return strings.Join([]string{"CreateOrderFormRequest", string(data)}, " ")
}
