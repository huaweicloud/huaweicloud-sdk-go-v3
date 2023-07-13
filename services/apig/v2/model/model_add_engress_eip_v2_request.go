package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddEngressEipV2Request Request Object
type AddEngressEipV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	Body *OpenEngressEipReq `json:"body,omitempty"`
}

func (o AddEngressEipV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddEngressEipV2Request struct{}"
	}

	return strings.Join([]string{"AddEngressEipV2Request", string(data)}, " ")
}
