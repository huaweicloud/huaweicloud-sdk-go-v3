package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartInstanceRequest Request Object
type StartInstanceRequest struct {

	// 实例id
	InstanceId string `json:"instance_id"`

	Body *StartInstanceParam `json:"body,omitempty"`
}

func (o StartInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartInstanceRequest struct{}"
	}

	return strings.Join([]string{"StartInstanceRequest", string(data)}, " ")
}
