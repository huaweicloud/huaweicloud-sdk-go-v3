package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddEipV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	Body *EipBindReq `json:"body,omitempty"`
}

func (o AddEipV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddEipV2Request struct{}"
	}

	return strings.Join([]string{"AddEipV2Request", string(data)}, " ")
}
