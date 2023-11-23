package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteReadonlyNodeRequest Request Object
type DeleteReadonlyNodeRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *DeleteReadonlyNodeRequestBody `json:"body,omitempty"`
}

func (o DeleteReadonlyNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteReadonlyNodeRequest struct{}"
	}

	return strings.Join([]string{"DeleteReadonlyNodeRequest", string(data)}, " ")
}
