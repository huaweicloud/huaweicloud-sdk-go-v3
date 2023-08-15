package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCaResourceInstancesRequest Request Object
type ListCaResourceInstancesRequest struct {
	Body *ListResourceInstancesRequestBody `json:"body,omitempty"`
}

func (o ListCaResourceInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCaResourceInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListCaResourceInstancesRequest", string(data)}, " ")
}
