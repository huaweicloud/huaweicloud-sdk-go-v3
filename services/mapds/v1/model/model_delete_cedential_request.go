package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteCedentialRequest struct {

	// 待删除的key的id。
	Clientid string `json:"clientid"`

	// IAM用户的token，无需特殊权限。
	XAuthToken string `json:"X-Auth-Token"`

	// 该字段填为“application/json;charset=utf8”。
	ContentType string `json:"Content-Type"`
}

func (o DeleteCedentialRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCedentialRequest struct{}"
	}

	return strings.Join([]string{"DeleteCedentialRequest", string(data)}, " ")
}
