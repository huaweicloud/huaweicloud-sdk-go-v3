package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchAssociateNaToNodesRequest struct {

	// 北向数据接收端点ID
	NaId string `json:"na_id" xml:"na_id"`

	// 批量删除delete，批量添加add
	Action string `json:"action" xml:"action"`

	Body *AuthorizeNa2NodesRequestDto `json:"body,omitempty" xml:"body"`
}

func (o BatchAssociateNaToNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAssociateNaToNodesRequest struct{}"
	}

	return strings.Join([]string{"BatchAssociateNaToNodesRequest", string(data)}, " ")
}
