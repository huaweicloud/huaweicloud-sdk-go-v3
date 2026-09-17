package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRdsInstanceConfigurationsNewRequest Request Object
type ListRdsInstanceConfigurationsNewRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ListRdsInstanceConfigurationsNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRdsInstanceConfigurationsNewRequest struct{}"
	}

	return strings.Join([]string{"ListRdsInstanceConfigurationsNewRequest", string(data)}, " ")
}
