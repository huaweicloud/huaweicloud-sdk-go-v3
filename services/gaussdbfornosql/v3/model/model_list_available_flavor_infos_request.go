package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAvailableFlavorInfosRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ListAvailableFlavorInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableFlavorInfosRequest struct{}"
	}

	return strings.Join([]string{"ListAvailableFlavorInfosRequest", string(data)}, " ")
}
