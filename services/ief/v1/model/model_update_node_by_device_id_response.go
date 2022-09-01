package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateNodeByDeviceIdResponse struct {
	UpdateNodes    *NodeUpdateByDevice `json:"update_nodes,omitempty" xml:"update_nodes"`
	HttpStatusCode int                 `json:"-"`
}

func (o UpdateNodeByDeviceIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNodeByDeviceIdResponse struct{}"
	}

	return strings.Join([]string{"UpdateNodeByDeviceIdResponse", string(data)}, " ")
}
