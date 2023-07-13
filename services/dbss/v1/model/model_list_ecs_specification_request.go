package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEcsSpecificationRequest Request Object
type ListEcsSpecificationRequest struct {
}

func (o ListEcsSpecificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEcsSpecificationRequest struct{}"
	}

	return strings.Join([]string{"ListEcsSpecificationRequest", string(data)}, " ")
}
