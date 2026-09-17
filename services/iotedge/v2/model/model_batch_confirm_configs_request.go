package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchConfirmConfigsRequest Request Object
type BatchConfirmConfigsRequest struct {

	// 边缘节点ID
	NodeId string `json:"node_id"`

	// 边侧第三方应用的模块ID
	IaId string `json:"ia_id"`

	// confirm
	Action string `json:"action"`

	Body *ConfirmIaConfigsRequestBody `json:"body,omitempty"`
}

func (o BatchConfirmConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchConfirmConfigsRequest struct{}"
	}

	return strings.Join([]string{"BatchConfirmConfigsRequest", string(data)}, " ")
}
