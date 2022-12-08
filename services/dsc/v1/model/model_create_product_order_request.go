package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
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
