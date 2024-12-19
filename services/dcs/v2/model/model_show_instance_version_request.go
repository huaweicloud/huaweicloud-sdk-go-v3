package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceVersionRequest Request Object
type ShowInstanceVersionRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceVersionRequest", string(data)}, " ")
}
