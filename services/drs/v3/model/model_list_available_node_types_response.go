package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailableNodeTypesResponse Response Object
type ListAvailableNodeTypesResponse struct {

	// Node规格列表
	NodeTypes      *[]QuerySupportNodeTypeBean `json:"node_types,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListAvailableNodeTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableNodeTypesResponse struct{}"
	}

	return strings.Join([]string{"ListAvailableNodeTypesResponse", string(data)}, " ")
}
