package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListObsBucketsRequest Request Object
type ListObsBucketsRequest struct {
}

func (o ListObsBucketsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListObsBucketsRequest struct{}"
	}

	return strings.Join([]string{"ListObsBucketsRequest", string(data)}, " ")
}
