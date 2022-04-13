package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteAppV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 应用编号

	AppId string `json:"app_id"`
}

func (o DeleteAppV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppV2Request struct{}"
	}

	return strings.Join([]string{"DeleteAppV2Request", string(data)}, " ")
}
