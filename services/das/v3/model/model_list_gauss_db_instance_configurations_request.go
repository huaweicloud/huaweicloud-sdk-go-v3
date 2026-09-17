package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGaussDbInstanceConfigurationsRequest Request Object
type ListGaussDbInstanceConfigurationsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ListGaussDbInstanceConfigurationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGaussDbInstanceConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"ListGaussDbInstanceConfigurationsRequest", string(data)}, " ")
}
