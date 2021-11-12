package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchAssociateNaToNodesResponse struct {
	// 授权北向NA信息到边缘节点列表的返回结构体，仅返回本次授权的节点列表信息

	Nodes          *[]QueryAuthorizedNodeDto `json:"nodes,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o BatchAssociateNaToNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAssociateNaToNodesResponse struct{}"
	}

	return strings.Join([]string{"BatchAssociateNaToNodesResponse", string(data)}, " ")
}
