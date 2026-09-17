package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesApiRequest Request Object
type ListInstancesApiRequest struct {
	Body *ListInstancesRequestBody `json:"body,omitempty"`
}

func (o ListInstancesApiRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesApiRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesApiRequest", string(data)}, " ")
}
