package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNodesResponse Response Object
type ShowNodesResponse struct {

	// 满足条件的设备总数
	Count *int32 `json:"count,omitempty"`

	Nodes          *[]NodeResponse `json:"nodes,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodesResponse struct{}"
	}

	return strings.Join([]string{"ShowNodesResponse", string(data)}, " ")
}
