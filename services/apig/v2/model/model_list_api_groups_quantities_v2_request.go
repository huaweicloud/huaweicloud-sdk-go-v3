package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListApiGroupsQuantitiesV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
}

func (o ListApiGroupsQuantitiesV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiGroupsQuantitiesV2Request struct{}"
	}

	return strings.Join([]string{"ListApiGroupsQuantitiesV2Request", string(data)}, " ")
}
