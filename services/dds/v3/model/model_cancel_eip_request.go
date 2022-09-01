package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CancelEipRequest struct {

	// 节点ID。
	NodeId string `json:"node_id" xml:"node_id"`
}

func (o CancelEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelEipRequest struct{}"
	}

	return strings.Join([]string{"CancelEipRequest", string(data)}, " ")
}
