package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCertResourceInstancesRequest Request Object
type ListCertResourceInstancesRequest struct {
	Body *ListResourceInstancesRequestBody `json:"body,omitempty"`
}

func (o ListCertResourceInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertResourceInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListCertResourceInstancesRequest", string(data)}, " ")
}
