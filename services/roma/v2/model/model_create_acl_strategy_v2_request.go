package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateAclStrategyV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *ApiAclCreate `json:"body,omitempty" xml:"body"`
}

func (o CreateAclStrategyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAclStrategyV2Request struct{}"
	}

	return strings.Join([]string{"CreateAclStrategyV2Request", string(data)}, " ")
}
