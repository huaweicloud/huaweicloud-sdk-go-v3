package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateApiV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	Body *ApiCreate `json:"body,omitempty"`
}

func (o CreateApiV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApiV2Request struct{}"
	}

	return strings.Join([]string{"CreateApiV2Request", string(data)}, " ")
}
