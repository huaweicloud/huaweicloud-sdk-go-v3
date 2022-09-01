package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateGatewayResponseV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 分组的编号
	GroupId string `json:"group_id" xml:"group_id"`

	// 响应编号
	ResponseId string `json:"response_id" xml:"response_id"`

	Body *ResponsesCreate `json:"body,omitempty" xml:"body"`
}

func (o UpdateGatewayResponseV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGatewayResponseV2Request struct{}"
	}

	return strings.Join([]string{"UpdateGatewayResponseV2Request", string(data)}, " ")
}
