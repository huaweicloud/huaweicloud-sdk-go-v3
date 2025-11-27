package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourceViewsRequest Request Object
type CreateResourceViewsRequest struct {
	Body *CreateResourceViewsRequestBody `json:"body,omitempty"`
}

func (o CreateResourceViewsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceViewsRequest struct{}"
	}

	return strings.Join([]string{"CreateResourceViewsRequest", string(data)}, " ")
}
