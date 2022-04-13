package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowBlockchainNodesResponse struct {
	// key:组织名，value：组织详细信息

	NodeOrgs       map[string]Org `json:"node_orgs,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowBlockchainNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBlockchainNodesResponse struct{}"
	}

	return strings.Join([]string{"ShowBlockchainNodesResponse", string(data)}, " ")
}
