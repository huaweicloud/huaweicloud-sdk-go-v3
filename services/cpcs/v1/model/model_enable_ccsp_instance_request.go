package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableCcspInstanceRequest Request Object
type EnableCcspInstanceRequest struct {

	// 密码服务实例ID
	InstanceId string `json:"instance_id"`
}

func (o EnableCcspInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableCcspInstanceRequest struct{}"
	}

	return strings.Join([]string{"EnableCcspInstanceRequest", string(data)}, " ")
}
