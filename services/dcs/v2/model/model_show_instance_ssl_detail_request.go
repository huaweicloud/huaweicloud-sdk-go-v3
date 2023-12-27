package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceSslDetailRequest Request Object
type ShowInstanceSslDetailRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceSslDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceSslDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceSslDetailRequest", string(data)}, " ")
}
