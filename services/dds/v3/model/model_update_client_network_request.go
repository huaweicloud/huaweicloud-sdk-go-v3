package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateClientNetworkRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *ClientNetworkRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateClientNetworkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClientNetworkRequest struct{}"
	}

	return strings.Join([]string{"UpdateClientNetworkRequest", string(data)}, " ")
}
