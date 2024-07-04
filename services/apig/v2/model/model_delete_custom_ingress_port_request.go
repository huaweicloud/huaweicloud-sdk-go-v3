package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCustomIngressPortRequest Request Object
type DeleteCustomIngressPortRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 实例自定义入方向端口ID。
	IngressPortId string `json:"ingress_port_id"`
}

func (o DeleteCustomIngressPortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCustomIngressPortRequest struct{}"
	}

	return strings.Join([]string{"DeleteCustomIngressPortRequest", string(data)}, " ")
}
