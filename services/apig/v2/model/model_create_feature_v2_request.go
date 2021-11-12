package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateFeatureV2Request struct {
	// 实例编号

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
