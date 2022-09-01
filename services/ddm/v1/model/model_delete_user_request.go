package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteUserRequest struct {

	// DDM实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 要删除的DDM帐号名称。
	Username string `json:"username" xml:"username"`
}

func (o DeleteUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteUserRequest struct{}"
	}

	return strings.Join([]string{"DeleteUserRequest", string(data)}, " ")
}
