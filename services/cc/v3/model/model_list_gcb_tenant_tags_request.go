package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGcbTenantTagsRequest Request Object
type ListGcbTenantTagsRequest struct {
}

func (o ListGcbTenantTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGcbTenantTagsRequest struct{}"
	}

	return strings.Join([]string{"ListGcbTenantTagsRequest", string(data)}, " ")
}
