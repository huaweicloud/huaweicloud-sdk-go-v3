package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteResourceGroupsRequest Request Object
type BatchDeleteResourceGroupsRequest struct {
	Body *BatchDeleteResourceGroupsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteResourceGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteResourceGroupsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteResourceGroupsRequest", string(data)}, " ")
}
