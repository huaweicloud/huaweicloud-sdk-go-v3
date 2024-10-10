package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDDoSAttackEventRequest Request Object
type ListDDoSAttackEventRequest struct {

	// 实例id
	InstanceId string `json:"instance_id"`

	Body *ListDDoSAttackEventRequestBody `json:"body,omitempty"`
}

func (o ListDDoSAttackEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDDoSAttackEventRequest struct{}"
	}

	return strings.Join([]string{"ListDDoSAttackEventRequest", string(data)}, " ")
}
