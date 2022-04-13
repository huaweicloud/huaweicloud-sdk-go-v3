package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListTagsOfTenantRequest struct {
}

func (o ListTagsOfTenantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsOfTenantRequest struct{}"
	}

	return strings.Join([]string{"ListTagsOfTenantRequest", string(data)}, " ")
}
