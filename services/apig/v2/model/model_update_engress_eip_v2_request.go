package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEngressEipV2Request Request Object
type UpdateEngressEipV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	Body *OpenEngressEipReq `json:"body,omitempty"`
}

func (o UpdateEngressEipV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEngressEipV2Request struct{}"
	}

	return strings.Join([]string{"UpdateEngressEipV2Request", string(data)}, " ")
}
