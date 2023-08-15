package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesByResourceTagsRequest Request Object
type ListInstancesByResourceTagsRequest struct {
	Body *ListInstancesByTagsRequestBody `json:"body,omitempty"`
}

func (o ListInstancesByResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesByResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesByResourceTagsRequest", string(data)}, " ")
}
