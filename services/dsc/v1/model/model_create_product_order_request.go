package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProductOrderRequest Request Object
type CreateProductOrderRequest struct {
	Body *PeriodOrderRequest `json:"body,omitempty"`
}

func (o CreateProductOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProductOrderRequest struct{}"
	}

	return strings.Join([]string{"CreateProductOrderRequest", string(data)}, " ")
}
