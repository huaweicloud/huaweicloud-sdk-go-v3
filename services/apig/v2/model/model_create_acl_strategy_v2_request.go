package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAclStrategyV2Request Request Object
type CreateAclStrategyV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	Body *ApiAclCreate `json:"body,omitempty"`
}

func (o CreateAclStrategyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAclStrategyV2Request struct{}"
	}

	return strings.Join([]string{"CreateAclStrategyV2Request", string(data)}, " ")
}
