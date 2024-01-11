package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEcnWithVpcRequest Request Object
type UpdateEcnWithVpcRequest struct {

	// 企业连接网络ID
	EcnId string `json:"ecn_id"`

	// 企业连接网络绑定关系ID
	RelationId string `json:"relation_id"`

	Body *UpdateEcnWithVpcRequestBody `json:"body,omitempty"`
}

func (o UpdateEcnWithVpcRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEcnWithVpcRequest struct{}"
	}

	return strings.Join([]string{"UpdateEcnWithVpcRequest", string(data)}, " ")
}
