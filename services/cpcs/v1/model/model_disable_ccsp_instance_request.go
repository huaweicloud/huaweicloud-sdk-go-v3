package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableCcspInstanceRequest Request Object
type DisableCcspInstanceRequest struct {

	// 密码服务实例ID
	InstanceId string `json:"instance_id"`
}

func (o DisableCcspInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableCcspInstanceRequest struct{}"
	}

	return strings.Join([]string{"DisableCcspInstanceRequest", string(data)}, " ")
}
