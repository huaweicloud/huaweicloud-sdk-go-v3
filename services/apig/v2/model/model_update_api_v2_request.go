package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateApiV2Request Request Object
type UpdateApiV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// API的编号
	ApiId string `json:"api_id"`

	Body *ApiCreate `json:"body,omitempty"`
}

func (o UpdateApiV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApiV2Request struct{}"
	}

	return strings.Join([]string{"UpdateApiV2Request", string(data)}, " ")
}
