package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSpecsRequest Request Object
type ListSpecsRequest struct {
}

func (o ListSpecsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSpecsRequest struct{}"
	}

	return strings.Join([]string{"ListSpecsRequest", string(data)}, " ")
}
