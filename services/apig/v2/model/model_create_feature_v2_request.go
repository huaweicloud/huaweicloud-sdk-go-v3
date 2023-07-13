package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFeatureV2Request Request Object
type CreateFeatureV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	Body *FeatureToggle `json:"body,omitempty"`
}

func (o CreateFeatureV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFeatureV2Request struct{}"
	}

	return strings.Join([]string{"CreateFeatureV2Request", string(data)}, " ")
}
