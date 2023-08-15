package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveEngressEipV2Request Request Object
type RemoveEngressEipV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`
}

func (o RemoveEngressEipV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveEngressEipV2Request struct{}"
	}

	return strings.Join([]string{"RemoveEngressEipV2Request", string(data)}, " ")
}
