package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateLiveDataApiV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`

	Body *LdApiCreate `json:"body,omitempty"`
}

func (o CreateLiveDataApiV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLiveDataApiV2Request struct{}"
	}

	return strings.Join([]string{"CreateLiveDataApiV2Request", string(data)}, " ")
}
