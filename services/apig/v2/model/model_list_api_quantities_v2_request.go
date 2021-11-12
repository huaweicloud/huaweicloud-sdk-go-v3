package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListApiQuantitiesV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
}

func (o ListApiQuantitiesV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiQuantitiesV2Request struct{}"
	}

	return strings.Join([]string{"ListApiQuantitiesV2Request", string(data)}, " ")
}
