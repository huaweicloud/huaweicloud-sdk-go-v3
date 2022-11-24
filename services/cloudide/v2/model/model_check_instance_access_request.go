package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CheckInstanceAccessRequest struct {

	// CloudIDE实例id
	InstanceId string `json:"instance_id"`
}

func (o CheckInstanceAccessRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckInstanceAccessRequest struct{}"
	}

	return strings.Join([]string{"CheckInstanceAccessRequest", string(data)}, " ")
}
