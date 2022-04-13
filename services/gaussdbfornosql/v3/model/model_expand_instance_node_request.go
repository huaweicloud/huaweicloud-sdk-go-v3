package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExpandInstanceNodeRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *ExpandInstanceNodeRequestBody `json:"body,omitempty"`
}

func (o ExpandInstanceNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandInstanceNodeRequest struct{}"
	}

	return strings.Join([]string{"ExpandInstanceNodeRequest", string(data)}, " ")
}
