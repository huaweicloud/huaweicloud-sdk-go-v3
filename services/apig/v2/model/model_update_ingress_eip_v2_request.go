package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIngressEipV2Request Request Object
type UpdateIngressEipV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	Body *OpenIngressEipReq `json:"body,omitempty"`
}

func (o UpdateIngressEipV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIngressEipV2Request struct{}"
	}

	return strings.Join([]string{"UpdateIngressEipV2Request", string(data)}, " ")
}
