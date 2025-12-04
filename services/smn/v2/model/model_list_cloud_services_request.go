package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCloudServicesRequest Request Object
type ListCloudServicesRequest struct {
}

func (o ListCloudServicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCloudServicesRequest struct{}"
	}

	return strings.Join([]string{"ListCloudServicesRequest", string(data)}, " ")
}
