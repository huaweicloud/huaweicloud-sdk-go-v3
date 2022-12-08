package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEcsSpecificationResponse struct {

	// ecs规格集合
	Specification  *[]EcsSpecificationBean `json:"specification,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListEcsSpecificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEcsSpecificationResponse struct{}"
	}

	return strings.Join([]string{"ListEcsSpecificationResponse", string(data)}, " ")
}
