package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudConnectionTagsRequest Request Object
type ListCloudConnectionTagsRequest struct {
}

func (o ListCloudConnectionTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudConnectionTagsRequest struct{}"
	}

	return strings.Join([]string{"ListCloudConnectionTagsRequest", string(data)}, " ")
}
