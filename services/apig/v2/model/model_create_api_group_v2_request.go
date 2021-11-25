package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateApiGroupV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`

	Body *ApiGroupCreate `json:"body,omitempty"`
}

func (o CreateApiGroupV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApiGroupV2Request struct{}"
	}

	return strings.Join([]string{"CreateApiGroupV2Request", string(data)}, " ")
}
