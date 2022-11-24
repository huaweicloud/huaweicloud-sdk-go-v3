package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListRegionsRequest struct {
}

func (o ListRegionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRegionsRequest struct{}"
	}

	return strings.Join([]string{"ListRegionsRequest", string(data)}, " ")
}
