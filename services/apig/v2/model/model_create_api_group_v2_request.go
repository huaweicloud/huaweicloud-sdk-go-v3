package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApiGroupV2Request Request Object
type CreateApiGroupV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	Body *ApiGroupCreate `json:"body,omitempty"`
}

func (o CreateApiGroupV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApiGroupV2Request struct{}"
	}

	return strings.Join([]string{"CreateApiGroupV2Request", string(data)}, " ")
}
