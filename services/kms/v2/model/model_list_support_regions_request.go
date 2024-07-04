package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSupportRegionsRequest Request Object
type ListSupportRegionsRequest struct {
}

func (o ListSupportRegionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportRegionsRequest struct{}"
	}

	return strings.Join([]string{"ListSupportRegionsRequest", string(data)}, " ")
}
