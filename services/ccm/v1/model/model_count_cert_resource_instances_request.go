package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountCertResourceInstancesRequest Request Object
type CountCertResourceInstancesRequest struct {
	Body *ListResourceInstancesRequestBody `json:"body,omitempty"`
}

func (o CountCertResourceInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountCertResourceInstancesRequest struct{}"
	}

	return strings.Join([]string{"CountCertResourceInstancesRequest", string(data)}, " ")
}
