package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNodeByDeviceIdResponse Response Object
type UpdateNodeByDeviceIdResponse struct {
	UpdateNodes    *NodeUpdateByDevice `json:"update_nodes,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o UpdateNodeByDeviceIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNodeByDeviceIdResponse struct{}"
	}

	return strings.Join([]string{"UpdateNodeByDeviceIdResponse", string(data)}, " ")
}
