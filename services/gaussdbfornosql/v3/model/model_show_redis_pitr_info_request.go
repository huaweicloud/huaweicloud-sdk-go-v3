package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRedisPitrInfoRequest Request Object
type ShowRedisPitrInfoRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowRedisPitrInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRedisPitrInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowRedisPitrInfoRequest", string(data)}, " ")
}
