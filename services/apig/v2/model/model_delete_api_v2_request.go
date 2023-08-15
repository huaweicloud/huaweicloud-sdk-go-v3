package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteApiV2Request Request Object
type DeleteApiV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// API的编号
	ApiId string `json:"api_id"`
}

func (o DeleteApiV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApiV2Request struct{}"
	}

	return strings.Join([]string{"DeleteApiV2Request", string(data)}, " ")
}
