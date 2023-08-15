package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEncryptdataNodesResponse Response Object
type ListEncryptdataNodesResponse struct {

	// 绑定的边缘节点总数
	Count *int32 `json:"count,omitempty"`

	// 边缘节点列表
	Nodes          *[]EncryptDataNodes `json:"nodes,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListEncryptdataNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEncryptdataNodesResponse struct{}"
	}

	return strings.Join([]string{"ListEncryptdataNodesResponse", string(data)}, " ")
}
