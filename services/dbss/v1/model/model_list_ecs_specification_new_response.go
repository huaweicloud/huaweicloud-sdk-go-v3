package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEcsSpecificationNewResponse Response Object
type ListEcsSpecificationNewResponse struct {

	// ecs规格集合
	Specifications *[]EcsSpecificationBean `json:"specifications,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListEcsSpecificationNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEcsSpecificationNewResponse struct{}"
	}

	return strings.Join([]string{"ListEcsSpecificationNewResponse", string(data)}, " ")
}
