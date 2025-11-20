package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceByInstanceIdRequest Request Object
type ShowInstanceByInstanceIdRequest struct {

	// 实例id
	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceByInstanceIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceByInstanceIdRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceByInstanceIdRequest", string(data)}, " ")
}
