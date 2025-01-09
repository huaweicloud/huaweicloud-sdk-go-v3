package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchSslRequest Request Object
type SwitchSslRequest struct {

	// DDM实例ID。
	InstanceId string `json:"instance_id"`

	Body *SwitchSslOpenRequest `json:"body,omitempty"`
}

func (o SwitchSslRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSslRequest struct{}"
	}

	return strings.Join([]string{"SwitchSslRequest", string(data)}, " ")
}
