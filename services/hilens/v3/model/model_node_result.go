package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeResult 部署节点的返回结果
type NodeResult struct {

	// 部署的节点ID
	NodeId string `json:"node_id"`

	// 部署到该节点失败时，返回的错误信息
	ErrorMessage *string `json:"error_message,omitempty"`

	// 部署到该节点失败时，返回的错误码
	ErrorCode *string `json:"error_code,omitempty"`
}

func (o NodeResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeResult struct{}"
	}

	return strings.Join([]string{"NodeResult", string(data)}, " ")
}
