package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteGatewayResponseV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 分组的编号

	GroupId string `json:"group_id"`
	// 响应编号

	ResponseId string `json:"response_id"`
}

func (o DeleteGatewayResponseV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGatewayResponseV2Request struct{}"
	}

	return strings.Join([]string{"DeleteGatewayResponseV2Request", string(data)}, " ")
}
