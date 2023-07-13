package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStarRequest Request Object
type ListStarRequest struct {
}

func (o ListStarRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStarRequest struct{}"
	}

	return strings.Join([]string{"ListStarRequest", string(data)}, " ")
}
