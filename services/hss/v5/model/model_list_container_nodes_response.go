package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListContainerNodesResponse Response Object
type ListContainerNodesResponse struct {

	// **参数解释**: 容器节点总数 **取值范围**: 取值0-65535
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 容器节点列表 **取值范围**: 取值0-65535
	DataList       *[]ContainerNodeInfo `json:"data_list,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListContainerNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListContainerNodesResponse struct{}"
	}

	return strings.Join([]string{"ListContainerNodesResponse", string(data)}, " ")
}
