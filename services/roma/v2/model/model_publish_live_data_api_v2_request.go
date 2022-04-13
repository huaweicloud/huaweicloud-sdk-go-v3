package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type PublishLiveDataApiV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 后端API的编号

	LdApiId string `json:"ld_api_id"`

	Body *LdApiDeploy `json:"body,omitempty"`
}

func (o PublishLiveDataApiV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishLiveDataApiV2Request struct{}"
	}

	return strings.Join([]string{"PublishLiveDataApiV2Request", string(data)}, " ")
}
