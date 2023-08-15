package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddIngressEipV2Request Request Object
type AddIngressEipV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	Body *OpenIngressEipReq `json:"body,omitempty"`
}

func (o AddIngressEipV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddIngressEipV2Request struct{}"
	}

	return strings.Join([]string{"AddIngressEipV2Request", string(data)}, " ")
}
