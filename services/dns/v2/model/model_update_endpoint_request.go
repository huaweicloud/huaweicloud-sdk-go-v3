package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateEndpointRequest struct {

	// 待修改endpoint的ID。。
	EndpointId string `json:"endpoint_id"`

	Body *UpdateEndpointInfo `json:"body,omitempty"`
}

func (o UpdateEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointRequest struct{}"
	}

	return strings.Join([]string{"UpdateEndpointRequest", string(data)}, " ")
}
