package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachEipRequest Request Object
type AttachEipRequest struct {

	// 需要绑定公网IP的节点ID。集群实例选择mongos节点，副本集实例选择primary或者secondary节点。
	NodeId string `json:"node_id"`

	Body *AttachEipRequestBody `json:"body,omitempty"`
}

func (o AttachEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachEipRequest struct{}"
	}

	return strings.Join([]string{"AttachEipRequest", string(data)}, " ")
}
