package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApiGroupsQuantitiesV2Request Request Object
type ListApiGroupsQuantitiesV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`
}

func (o ListApiGroupsQuantitiesV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiGroupsQuantitiesV2Request struct{}"
	}

	return strings.Join([]string{"ListApiGroupsQuantitiesV2Request", string(data)}, " ")
}
