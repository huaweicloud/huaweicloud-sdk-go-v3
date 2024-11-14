package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplyConfigurationToInstancesRequestBody struct {

	// 实例ID列表对象。
	InstanceIds []string `json:"instance_ids"`
}

func (o ApplyConfigurationToInstancesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyConfigurationToInstancesRequestBody struct{}"
	}

	return strings.Join([]string{"ApplyConfigurationToInstancesRequestBody", string(data)}, " ")
}
