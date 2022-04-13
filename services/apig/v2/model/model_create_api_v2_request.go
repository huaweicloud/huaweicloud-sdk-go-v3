package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateApiV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`

	Body *ApiCreate `json:"body,omitempty"`
}

func (o CreateApiV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApiV2Request struct{}"
	}

	return strings.Join([]string{"CreateApiV2Request", string(data)}, " ")
}
