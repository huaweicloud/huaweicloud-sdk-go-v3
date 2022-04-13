package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListNodeTypesRequest struct {
}

func (o ListNodeTypesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNodeTypesRequest struct{}"
	}

	return strings.Join([]string{"ListNodeTypesRequest", string(data)}, " ")
}
