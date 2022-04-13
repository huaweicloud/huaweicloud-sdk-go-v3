package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateLiveDataApiV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 后端API的编号

	LdApiId string `json:"ld_api_id"`

	Body *LdApiCreate `json:"body,omitempty"`
}

func (o UpdateLiveDataApiV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLiveDataApiV2Request struct{}"
	}

	return strings.Join([]string{"UpdateLiveDataApiV2Request", string(data)}, " ")
}
