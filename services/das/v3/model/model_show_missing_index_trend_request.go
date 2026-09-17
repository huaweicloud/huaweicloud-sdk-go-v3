package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMissingIndexTrendRequest Request Object
type ShowMissingIndexTrendRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ShowMissingIndexTrendRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMissingIndexTrendRequest struct{}"
	}

	return strings.Join([]string{"ShowMissingIndexTrendRequest", string(data)}, " ")
}
