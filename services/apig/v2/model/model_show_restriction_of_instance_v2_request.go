package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRestrictionOfInstanceV2Request Request Object
type ShowRestrictionOfInstanceV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`
}

func (o ShowRestrictionOfInstanceV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRestrictionOfInstanceV2Request struct{}"
	}

	return strings.Join([]string{"ShowRestrictionOfInstanceV2Request", string(data)}, " ")
}
