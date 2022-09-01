package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowRestrictionOfInstanceV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`
}

func (o ShowRestrictionOfInstanceV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRestrictionOfInstanceV2Request struct{}"
	}

	return strings.Join([]string{"ShowRestrictionOfInstanceV2Request", string(data)}, " ")
}
