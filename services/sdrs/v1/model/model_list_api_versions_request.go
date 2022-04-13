package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListApiVersionsRequest struct {
}

func (o ListApiVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListApiVersionsRequest", string(data)}, " ")
}
