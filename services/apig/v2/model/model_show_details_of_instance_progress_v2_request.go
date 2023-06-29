package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDetailsOfInstanceProgressV2Request Request Object
type ShowDetailsOfInstanceProgressV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`
}

func (o ShowDetailsOfInstanceProgressV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfInstanceProgressV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfInstanceProgressV2Request", string(data)}, " ")
}
