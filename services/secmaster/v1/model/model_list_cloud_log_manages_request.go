package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudLogManagesRequest Request Object
type ListCloudLogManagesRequest struct {
}

func (o ListCloudLogManagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudLogManagesRequest struct{}"
	}

	return strings.Join([]string{"ListCloudLogManagesRequest", string(data)}, " ")
}
