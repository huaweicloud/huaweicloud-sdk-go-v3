package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGatewayResponseV2Request Request Object
type UpdateGatewayResponseV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 分组的编号
	GroupId string `json:"group_id"`

	// 响应编号
	ResponseId string `json:"response_id"`

	Body *ResponsesCreate `json:"body,omitempty"`
}

func (o UpdateGatewayResponseV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGatewayResponseV2Request struct{}"
	}

	return strings.Join([]string{"UpdateGatewayResponseV2Request", string(data)}, " ")
}
