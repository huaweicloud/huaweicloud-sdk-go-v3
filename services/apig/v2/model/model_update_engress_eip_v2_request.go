package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateEngressEipV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`

	Body *OpenEngressEipReq `json:"body,omitempty"`
}

func (o UpdateEngressEipV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEngressEipV2Request struct{}"
	}

	return strings.Join([]string{"UpdateEngressEipV2Request", string(data)}, " ")
}
