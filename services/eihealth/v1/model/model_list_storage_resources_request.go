package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStorageResourcesRequest Request Object
type ListStorageResourcesRequest struct {
}

func (o ListStorageResourcesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStorageResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListStorageResourcesRequest", string(data)}, " ")
}
