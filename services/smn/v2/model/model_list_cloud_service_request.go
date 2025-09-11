package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudServiceRequest Request Object
type ListCloudServiceRequest struct {
}

func (o ListCloudServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudServiceRequest struct{}"
	}

	return strings.Join([]string{"ListCloudServiceRequest", string(data)}, " ")
}
