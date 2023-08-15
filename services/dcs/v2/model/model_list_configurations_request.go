package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigurationsRequest Request Object
type ListConfigurationsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ListConfigurationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"ListConfigurationsRequest", string(data)}, " ")
}
