package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMongosNodeRequest Request Object
type DeleteMongosNodeRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *DeleteMongosNodeRequestBody `json:"body,omitempty"`
}

func (o DeleteMongosNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMongosNodeRequest struct{}"
	}

	return strings.Join([]string{"DeleteMongosNodeRequest", string(data)}, " ")
}
