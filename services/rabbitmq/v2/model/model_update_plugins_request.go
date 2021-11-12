package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePluginsRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *UpdatePluginsReq `json:"body,omitempty"`
}

func (o UpdatePluginsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePluginsRequest struct{}"
	}

	return strings.Join([]string{"UpdatePluginsRequest", string(data)}, " ")
}
