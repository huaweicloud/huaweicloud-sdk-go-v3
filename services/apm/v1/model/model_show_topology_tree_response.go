package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTopologyTreeResponse struct {
	TopologyTree   *TopologyTree `json:"topology_tree,omitempty" xml:"topology_tree"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowTopologyTreeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopologyTreeResponse struct{}"
	}

	return strings.Join([]string{"ShowTopologyTreeResponse", string(data)}, " ")
}
