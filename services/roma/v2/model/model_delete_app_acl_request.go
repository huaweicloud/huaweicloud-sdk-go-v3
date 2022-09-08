package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteAppAclRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 应用编号
	AppId string `json:"app_id"`
}

func (o DeleteAppAclRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppAclRequest struct{}"
	}

	return strings.Join([]string{"DeleteAppAclRequest", string(data)}, " ")
}
