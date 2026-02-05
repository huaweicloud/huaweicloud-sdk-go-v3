package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppV2Request Request Object
type DeleteAppV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 凭据编号
	AppId string `json:"app_id"`
}

func (o DeleteAppV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppV2Request struct{}"
	}

	return strings.Join([]string{"DeleteAppV2Request", string(data)}, " ")
}
